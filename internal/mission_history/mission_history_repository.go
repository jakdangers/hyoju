package mission_history

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
)

type missionHistoryRepository struct {
	gormDB *gorm.DB
}

func NewMissionHistoryRepository(gormDB *gorm.DB) *missionHistoryRepository {
	return &missionHistoryRepository{gormDB: gormDB}
}

var _ entity.MissionHistoryRepository = (*missionHistoryRepository)(nil)

func (m missionHistoryRepository) CreateMissionHistory(ctx context.Context, missionHistory *entity.MissionHistory) (*entity.MissionHistory, error) {
	//TODO implement me
	panic("implement me")
}
