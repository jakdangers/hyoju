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

func (m challengeHistoryRepository) ListGroupChallengeHistories(ctx context.Context, params entity.ListGroupChallengeHistoriesParams) ([]entity.ChallengeHistory, error) {
	const op cerrors.Op = "challengeHistory/repository/ListGroupChallengeHistories"

	var challengeHistories []entity.ChallengeHistory
	if result := m.gormDB.WithContext(ctx).
		Where("created_at >= ? AND created_at < ? AND challenge_id IN ?", params.StartDateTime, params.StartDateTime, params.ChallengeID).
		Find(&challengeHistories); result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return challengeHistories, nil
}
