package challenge

import (
	"context"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type challengeService struct {
	challengeRepo            entity.ChallengeRepository
	challengeParticipantRepo entity.MissionParticipantRepository
	userRepo                 entity.UserRepository
}

func NewChallengeService(missionRepo entity.ChallengeRepository, missionParticipantRepo entity.MissionParticipantRepository, userRepo entity.UserRepository) *challengeService {
	return &challengeService{
		challengeRepo:            missionRepo,
		challengeParticipantRepo: missionParticipantRepo,
		userRepo:                 userRepo,
	}
}

var _ entity.ChallengeService = (*challengeService)(nil)

func (m challengeService) CreateMission(ctx context.Context, req entity.CreateMissionRequest) (*entity.CreateMissionResponse, error) {
	const op cerrors.Op = "challenge/service/createMission"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return nil, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	//planTime := time.Duration(req.PlanTime.Hour())*time.Hour + time.Duration(req.PlanTime.Minute())*time.Minute

	mission, err := m.challengeRepo.CreateChallenge(ctx, &entity.Challenge{
		UserID:    userID,
		Title:     req.Title,
		Emoji:     req.Emoji,
		Duration:  req.Duration,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		PlanTime:  req.PlanTime,
		Alarm:     req.Alarm,
		WeekDay:   entity.ConvertDaysOfWeekToInt(req.WeekDay),
		Type:      entity.Single,
		Status:    entity.Active,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	_, err = m.challengeParticipantRepo.CreateMissionParticipant(ctx, &entity.MissionParticipant{
		UserID:    userID,
		MissionID: mission.ID,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &entity.CreateMissionResponse{
		MissionID: mission.ID,
	}, nil
}

func (m challengeService) GetMission(ctx context.Context, req entity.GetMissionRequest) (*entity.GetMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m challengeService) ListMissions(ctx context.Context, req entity.ListMissionsRequest) (*entity.ListMissionsResponse, error) {
	const op cerrors.Op = "challenge/service/listMissions"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return nil, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	missions, err := m.challengeRepo.ListChallenges(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var missionDTOs []entity.MissionDTO
	for _, mission := range missions {
		missionDTOs = append(missionDTOs, entity.MissionDTOFrom(mission))
	}

	return &entity.ListMissionsResponse{
		Missions: missionDTOs,
	}, nil
}

func (m challengeService) PatchMission(ctx context.Context, req entity.PatchMissionRequest) (*entity.PatchMissionResponse, error) {
	const op cerrors.Op = "challenge/service/patchMission"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return nil, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	mission, err := m.challengeRepo.GetChallenge(ctx, req.ID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	if req.Title != nil {
		mission.Title = *req.Title
	}
	if req.Emoji != nil {
		mission.Emoji = *req.Emoji
	}
	if req.Duration != nil {
		mission.Duration = *req.Duration
	}
	if req.StartDate != nil {
		mission.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		mission.EndDate = *req.EndDate
	}
	if req.PlanTime != nil {
		//duration := time.Duration(req.PlanTime.Hour())*time.Hour + time.Duration(req.PlanTime.Minute())*time.Minute
		mission.PlanTime = *req.PlanTime
	}
	if req.Alarm != nil {
		mission.Alarm = *req.Alarm
	}
	if req.WeekDay != nil {
		mission.WeekDay = entity.ConvertDaysOfWeekToInt(req.WeekDay)
	}
	if req.Type != nil {
		mission.Type = *req.Type
	}
	if req.Status != nil {
		mission.Status = entity.ChallengeStatus(*req.Status)
	}

	patchMission, err := m.challengeRepo.PatchChallenge(ctx, mission)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &entity.PatchMissionResponse{MissionDTO: entity.MissionDTOFrom(*patchMission)}, nil
}
