package group_challenge

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type groupChallengeRepository struct {
	gormDB *gorm.DB
}

func NewGroupChallengeRepository(gormDB *gorm.DB) *groupChallengeRepository {
	return &groupChallengeRepository{gormDB: gormDB}
}

var _ entity.GroupChallengeRepository = (*groupChallengeRepository)(nil)

func (g groupChallengeRepository) CreateGroupChallenge(ctx context.Context, groupChallenge *entity.GroupChallenge) (*entity.GroupChallenge, error) {
	var op cerrors.Op = "groupChallenge/repository/createGroupChallenge"

	result := g.gormDB.WithContext(ctx).Create(groupChallenge)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return groupChallenge, nil
}

func (g groupChallengeRepository) ListGroupChallenges(c context.Context, req entity.ListGroupChallengesParams) (entity.GroupChallenges, error) {
	//TODO implement me
	panic("implement me")
}
