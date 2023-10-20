package mission

import (
	"context"
	"fmt"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type missionService struct {
	missionRepo            entity.MissionRepository
	missionParticipantRepo entity.MissionParticipantRepository
	userRepo               entity.UserRepository
}

func NewMissionService(missionRepo entity.MissionRepository, missionParticipantRepo entity.MissionParticipantRepository, userRepo entity.UserRepository) *missionService {
	return &missionService{
		missionRepo:            missionRepo,
		missionParticipantRepo: missionParticipantRepo,
		userRepo:               userRepo,
	}
}

var _ entity.MissionService = (*missionService)(nil)

func (m missionService) CreateMission(ctx context.Context, req entity.CreateMissionRequest) (*entity.CreateMissionResponse, error) {
	const op cerrors.Op = "mission/service/createMission"

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

	fmt.Println(req.StartDate.UTC())
	fmt.Println(req.EndDate.UTC())

	mission, err := m.missionRepo.CreateMission(ctx, &entity.Mission{
		AuthorID:  userID,
		Title:     req.Title,
		Emoji:     req.Emoji,
		Duration:  req.Duration,
		StartDate: req.StartDate.UTC(),
		EndDate:   req.EndDate.UTC(),
		PlanTime:  req.PlanTime.UTC(),
		Alarm:     req.Alarm,
		WeekDay:   entity.ConvertDaysOfWeekToInt(req.WeekDay),
		Type:      entity.Single,
		Status:    entity.Active,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	_, err = m.missionParticipantRepo.CreateMissionParticipant(ctx, &entity.MissionParticipant{
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

func (m missionService) GetMission(ctx context.Context, req entity.GetMissionRequest) (*entity.GetMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m missionService) ListMissions(ctx context.Context, req entity.ListMissionsRequest) (*entity.ListMissionsResponse, error) {
	const op cerrors.Op = "mission/service/listMissions"

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

	missions, err := m.missionRepo.ListMissions(ctx, userID)
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

func (m missionService) PatchMission(ctx context.Context, req entity.PatchMissionRequest) (*entity.PatchMissionResponse, error) {
	const op cerrors.Op = "mission/service/patchMission"

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

	mission, err := m.missionRepo.GetMission(ctx, req.ID)
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
		mission.Status = *req.Status
	}

	patchMission, err := m.missionRepo.PatchMission(ctx, mission)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &entity.PatchMissionResponse{MissionDTO: entity.MissionDTOFrom(*patchMission)}, nil
}
