package mission_history

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
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

func (m missionHistoryRepository) ListMultipleModeMissionHistories(ctx context.Context, params entity.ListMultipleMissionHistoriesParams) ([]entity.MissionHistory, error) {
	const op cerrors.Op = "missionHistory/repository/ListMultipleModeMissionHistories"

	var missionHistories []entity.MissionHistory
	if result := m.gormDB.WithContext(ctx).
		Where("user_id = ? AND mission_id IN ?", params.UserID, params.MissionIDs).
		Find(&missionHistories); result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return missionHistories, nil
}
