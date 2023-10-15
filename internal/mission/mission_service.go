package mission

import (
	"context"
	"pixelix/dto"
	"pixelix/entity"
)

type missionService struct {
	repository entity.MissionRepository
}

func NewMissionService(repo entity.MissionRepository) *missionService {
	return &missionService{
		repository: repo,
	}
}

var _ entity.MissionService = (*missionService)(nil)

func (m missionService) CreateMission(ctx context.Context, req dto.CreateMissionRequest) (dto.CreateMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}
