package challenge

import (
	"context"
	"pixelix/entity"
)

type challengeService struct {
	userRepo      entity.UserRepository
	challengeRepo entity.ChallengeRepository
}

func NewChallengeService(userRepo entity.UserRepository, challengeRepo entity.ChallengeRepository) *challengeService {
	return &challengeService{
		userRepo:      userRepo,
		challengeRepo: challengeRepo,
	}
}

var _ entity.ChallengeService = (*challengeService)(nil)

func (cs challengeService) CreateChallenge(c context.Context, req entity.CreateChallengeRequest) (*entity.CreateChallengeResponse, error) {
	//TODO implement me
	panic("implement me")
}
