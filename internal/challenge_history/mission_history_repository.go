package challenge_history

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type challengeHistoryRepository struct {
	gormDB *gorm.DB
}

func NewChallengeHistoryRepository(gormDB *gorm.DB) *challengeHistoryRepository {
	return &challengeHistoryRepository{gormDB: gormDB}
}

var _ entity.ChallengeHistoryRepository = (*challengeHistoryRepository)(nil)

func (m challengeHistoryRepository) CreateChallengeHistory(ctx context.Context, missionHistory *entity.ChallengeHistory) (*entity.ChallengeHistory, error) {
	//TODO implement me
	panic("implement me")
}

func (m challengeHistoryRepository) ListMultiChallengeHistories(ctx context.Context, params entity.ListMultipleMissionHistoriesParams) ([]entity.ChallengeHistory, error) {
	const op cerrors.Op = "challengeHistory/repository/ListMultiChallengeHistories"

	var challengeHistories []entity.ChallengeHistory
	if result := m.gormDB.WithContext(ctx).
		Where("user_id = ? AND challenge_id IN ?", params.UserID, params.ChallengeIDs).
		Find(&challengeHistories); result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return challengeHistories, nil
}
