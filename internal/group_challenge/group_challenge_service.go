package group_challenge

import (
	"context"
	"pixelix/entity"
)

type groupChallengeService struct {
	groupChallengeRepository entity.GroupChallengeRepository
}

func NewGroupChallengeService(groupChallengeRepository entity.GroupChallengeRepository) *groupChallengeService {
	return &groupChallengeService{groupChallengeRepository: groupChallengeRepository}
}

var _ entity.GroupChallengeService = (*groupChallengeService)(nil)

func (g groupChallengeService) CreateGroupChallenge(c context.Context, req entity.CreateGroupChallengeRequest) (entity.CreateGroupChallengeResponse, error) {
	//TODO implement me
	panic("implement me")
}
