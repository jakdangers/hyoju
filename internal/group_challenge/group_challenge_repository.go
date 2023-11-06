package group_challenge

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
)

type groupChallengeRepository struct {
	gormDB *gorm.DB
}

func NewGroupChallengeRepository(gormDB *gorm.DB) *groupChallengeRepository {
	return &groupChallengeRepository{gormDB: gormDB}
}

var _ entity.GroupChallengeRepository = (*groupChallengeRepository)(nil)

func (g groupChallengeRepository) CreateGroupChallenge(ctx context.Context, groupChallenge *entity.GroupChallenge) (*entity.GroupChallenge, error) {
	//TODO implement me
	panic("implement me")
}
