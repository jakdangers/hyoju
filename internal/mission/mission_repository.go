package mission

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type missionRepository struct {
	gormDB *gorm.DB
}

func NewMissionRepository(gormDB *gorm.DB) *missionRepository {
	return &missionRepository{
		gormDB: gormDB,
	}
}

var _ entity.MissionRepository = (*missionRepository)(nil)

func (m missionRepository) CreateMission(ctx context.Context, mission *entity.Mission) (*entity.Mission, error) {
	const op cerrors.Op = "mission/repository/createMission"

	result := m.gormDB.WithContext(ctx).Create(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}

func (m missionRepository) ListMissions(ctx context.Context, userID entity.BinaryUUID) ([]entity.Mission, error) {
	const op cerrors.Op = "mission/repository/listMissions"

	var missions []entity.Mission
	result := m.gormDB.WithContext(ctx).Where("author_id = ? AND status = ?", userID, entity.Active).Find(&missions)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return missions, nil
}
