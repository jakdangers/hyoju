package mission_participant

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type missionParticipantRepository struct {
	gormDB *gorm.DB
}

func NewMissionParticipantRepository(gormDB *gorm.DB) *missionParticipantRepository {
	return &missionParticipantRepository{gormDB: gormDB}
}

var _ entity.MissionParticipantRepository = (*missionParticipantRepository)(nil)

func (m missionParticipantRepository) CreateMissionParticipant(ctx context.Context, participant *entity.MissionParticipant) (*entity.MissionParticipant, error) {
	const op cerrors.Op = "missionParticipant/repository/createMissionParticipant"

	if err := m.gormDB.WithContext(ctx).Create(participant).Error; err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return participant, nil
}

func (m missionParticipantRepository) ListMissionParticipants(ctx context.Context, missionID uint) ([]entity.MissionParticipant, error) {
	const op cerrors.Op = "missionParticipant/repository/listMissionParticipants"

	var participants []entity.MissionParticipant
	if err := m.gormDB.WithContext(ctx).Find(&participants, "mission_id = ?", missionID).Error; err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return participants, nil
}
