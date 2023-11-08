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

func (g groupChallengeService) ListGroupChallenges(c context.Context, req entity.ListGroupChallengesRequest) (*entity.ListGroupChallengesResponse, error) {
	var op cerrors.Op = "groupChallenge/service/listGroupChallenges"

	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	groupChallenges, err := g.groupChallengeRepo.ListGroupChallenges(c, entity.ListGroupChallengesParams{
		UserID: userID,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var groupChallengeDtos []entity.GroupChallengeDto
	for _, groupChallenge := range groupChallenges {
		groupChallengeDtos = append(groupChallengeDtos, entity.GroupChallengeDtoFrom(groupChallenge))
	}

	return &entity.ListGroupChallengesResponse{
		GroupChallenges: groupChallengeDtos,
	}, nil
}
