package group_challenge

import (
	"context"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type groupChallengeService struct {
	groupChallengeRepo entity.GroupChallengeRepository
	userRepo           entity.UserRepository
}

func NewGroupChallengeService(groupChallengeRepo entity.GroupChallengeRepository, userRepo entity.UserRepository) *groupChallengeService {
	return &groupChallengeService{
		groupChallengeRepo: groupChallengeRepo,
		userRepo:           userRepo,
	}
}

var _ entity.GroupChallengeService = (*groupChallengeService)(nil)

func (g groupChallengeService) CreateGroupChallenge(c context.Context, req entity.CreateGroupChallengeRequest) error {
	var op cerrors.Op = "groupChallenge/service/createGroupChallenge"

	_, err := g.groupChallengeRepo.CreateGroupChallenge(c, &entity.GroupChallenge{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err)
	}

	return nil
}
