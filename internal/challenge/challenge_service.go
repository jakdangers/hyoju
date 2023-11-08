package challenge

import (
	"context"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type challengeService struct {
	challengeRepo            entity.ChallengeRepository
	challengeParticipantRepo entity.ChallengeParticipantRepository
	userRepo                 entity.UserRepository
}

func NewChallengeService(missionRepo entity.ChallengeRepository, missionParticipantRepo entity.ChallengeParticipantRepository, userRepo entity.UserRepository) *challengeService {
	return &challengeService{
		challengeRepo:            missionRepo,
		challengeParticipantRepo: missionParticipantRepo,
		userRepo:                 userRepo,
	}
}

var _ entity.ChallengeService = (*challengeService)(nil)

func (m challengeService) CreateChallenge(ctx context.Context, req entity.CreateChallengeRequest) (*entity.CreateChallengeResponse, error) {
	const op cerrors.Op = "challenge/service/createChallenge"

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

	code := entity.BinaryUUIDNew().ToCode()
	for {
		findChallenge, err := m.challengeRepo.ChallengeFindByCode(ctx, code)
		if err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err)
		}
		if findChallenge == nil {
			break
		}
		code = entity.BinaryUUIDNew().ToCode()
	}

	challenge, err := m.challengeRepo.CreateChallenge(ctx, &entity.Challenge{
		UserID:    userID,
		Title:     req.Title,
		Emoji:     req.Emoji,
		Duration:  req.Duration,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		PlanTime:  req.PlanTime,
		Alarm:     req.Alarm,
		WeekDay:   entity.ConvertDaysOfWeekToInt(req.WeekDay),
		Type:      req.Type,
		Status:    entity.ChallengeStatusActivate,
		Code:      code,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	_, err = m.challengeParticipantRepo.CreateChallengeParticipant(ctx, &entity.ChallengeParticipant{
		UserID:      userID,
		ChallengeID: challenge.ID,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &entity.CreateChallengeResponse{
		ChallengeID: challenge.ID,
	}, nil
}

func (m challengeService) GetChallenge(ctx context.Context, req entity.GetChallengeRequest) (*entity.GetChallengeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m challengeService) ListChallenges(ctx context.Context, req entity.ListChallengesRequest) (*entity.ListChallengesResponse, error) {
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

	challenges, err := m.challengeRepo.ListChallenges(ctx, entity.ListChallengesParams{
		UserID: userID,
		Type:   req.Type,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var challengeDTOS []entity.ChallengeDto
	for _, challenge := range challenges {
		challengeDTOS = append(challengeDTOS, entity.ChallengeDtoFrom(challenge))
	}

	return &entity.ListChallengesResponse{
		Challenges: challengeDTOS,
	}, nil
}

func (m challengeService) PatchChallenge(ctx context.Context, req entity.PatchChallengeRequest) (*entity.PatchChallengeResponse, error) {
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

	patchMission, err := m.challengeRepo.PatchChallenge(ctx, mission)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &entity.PatchChallengeResponse{ChallengeDto: entity.ChallengeDtoFrom(*patchMission)}, nil
}
