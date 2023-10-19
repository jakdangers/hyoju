package single_mode

import (
	"context"
	"pixelix/entity"
)

type singleModeService struct {
	missionRepo            entity.MissionRepository
	missionParticipantRepo entity.MissionParticipantRepository
	missionHistoryRepo     entity.MissionHistoryRepository
}

func NewSingleModeService(missionRepo entity.MissionRepository, missionParticipantRepo entity.MissionParticipantRepository, missionHistoryRepo entity.MissionHistoryRepository) *singleModeService {
	return &singleModeService{
		missionRepo:            missionRepo,
		missionParticipantRepo: missionParticipantRepo,
		missionHistoryRepo:     missionHistoryRepo,
	}
}

var _ entity.SingleModeService = (*singleModeService)(nil)

func (s singleModeService) CreateSingleModeMissionHistories(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
