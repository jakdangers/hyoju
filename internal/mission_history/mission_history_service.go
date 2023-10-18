package mission_history

import (
	"context"
	"pixelix/entity"
)

type missionHistoryService struct {
	missionRepo            entity.MissionRepository
	missionParticipantRepo entity.MissionParticipantRepository
	userRepo               entity.UserRepository
}

func NewMissionHistoryService(missionRepo entity.MissionRepository, missionParticipantRepo entity.MissionParticipantRepository, userRepo entity.UserRepository) *missionHistoryService {
	return &missionHistoryService{
		missionRepo:            missionRepo,
		missionParticipantRepo: missionParticipantRepo,
		userRepo:               userRepo,
	}
}

var _ entity.MissionHistoryService = (*missionHistoryService)(nil)

func (m missionHistoryService) CreateMissionHistory(ctx context.Context, req entity.CreateMissionHistoryRequest) (*entity.CreateMissionHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
