package mission_history

import (
	"context"
	"fmt"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type missionHistoryService struct {
	missionRepo            entity.MissionRepository
	missionParticipantRepo entity.MissionParticipantRepository
	missionHistoryRepo     entity.MissionHistoryRepository
	userRepo               entity.UserRepository
}

func NewMissionHistoryService(missionRepo entity.MissionRepository, missionParticipantRepo entity.MissionParticipantRepository, missionHistoryRepo entity.MissionHistoryRepository, userRepo entity.UserRepository) *missionHistoryService {
	return &missionHistoryService{
		missionRepo:            missionRepo,
		missionParticipantRepo: missionParticipantRepo,
		missionHistoryRepo:     missionHistoryRepo,
		userRepo:               userRepo,
	}
}

var _ entity.MissionHistoryService = (*missionHistoryService)(nil)

func (m missionHistoryService) CreateMissionHistory(ctx context.Context, req entity.CreateMissionHistoryRequest) (*entity.CreateMissionHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m missionHistoryService) ListMultipleMissionHistories(ctx context.Context, req entity.ListMissionHistoriesRequest) (*entity.ListMissionHistoriesResponse, error) {
	const op cerrors.Op = "missionHistory/service/ListMultipleModeMissionHistories"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	_, err = m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	missions, err := m.missionRepo.ListMultipleModeMissions(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var missionIDs []uint
	for _, mission := range missions {
		missionIDs = append(missionIDs, mission.ID)
	}

	histories, err := m.missionHistoryRepo.ListMultipleModeMissionHistories(ctx, entity.ListMultipleMissionHistoriesParams{
		UserID:     userID,
		MissionIDs: missionIDs,
	})

	var historyIDs []uint
	for _, history := range histories {
		historyIDs = append(historyIDs, history.ID)
	}

	for _, mission := range missions {

	}

	fmt.Println(histories)

	return nil, nil
}
