package mission

import (
	"context"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type missionService struct {
	missionRepo entity.MissionRepository
	userRepo    entity.UserRepository
}

func NewMissionService(missionRepo entity.MissionRepository, userRepo entity.UserRepository) *missionService {
	return &missionService{
		missionRepo: missionRepo,
		userRepo:    userRepo,
	}
}

var _ entity.MissionService = (*missionService)(nil)

func (m missionService) CreateMission(ctx context.Context, req entity.CreateMissionRequest) (entity.CreateMissionResponse, error) {
	const op cerrors.Op = "mission/service/createMission"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return entity.CreateMissionResponse{}, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return entity.CreateMissionResponse{}, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return entity.CreateMissionResponse{}, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	days, err := entity.ConvertDaysOfWeekToInt(req.WeekDay)
	if err != nil {
		return entity.CreateMissionResponse{}, cerrors.E(op, cerrors.Invalid, err)
	}

	mission, err := m.missionRepo.CreateMission(ctx, &entity.Mission{
		AuthorID:  userID,
		Title:     req.Title,
		Emoji:     req.Emoji,
		Duration:  req.Duration,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		PlanTime:  req.PlanTime,
		Alarm:     req.Alarm,
		WeekDay:   days,
		Type:      entity.Single,
		Status:    entity.Active,
	})
	if err != nil {
		return entity.CreateMissionResponse{}, cerrors.E(op, cerrors.Internal, err)
	}

	return entity.CreateMissionResponse{
		ID: mission.ID,
	}, nil
}

func (m missionService) PatchMission(ctx context.Context, req entity.PatchMissionRequest) (entity.PatchMissionResponse, error) {
	const op cerrors.Op = "mission/service/patchMission"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return entity.PatchMissionResponse{}, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return entity.PatchMissionResponse{}, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return entity.PatchMissionResponse{}, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	return entity.PatchMissionResponse{}, nil
}

func (m missionService) ListMissions(ctx context.Context, req entity.ListMissionsRequest) (entity.ListMissionsResponse, error) {
	const op cerrors.Op = "mission/service/listMissions"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return entity.ListMissionsResponse{}, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return entity.ListMissionsResponse{}, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return entity.ListMissionsResponse{}, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	missions, err := m.missionRepo.ListMissions(ctx, userID)
	if err != nil {
		return entity.ListMissionsResponse{}, cerrors.E(op, cerrors.Internal, err)
	}

	var missionDTOs []entity.MissionDTO
	for _, mission := range missions {
		missionDTOs = append(missionDTOs, entity.MissionDTOFrom(mission))
	}

	return entity.ListMissionsResponse{
		Missions: missionDTOs,
	}, nil
}
