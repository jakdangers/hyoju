package challenge

import (
	"context"
	"gorm.io/gorm"
	"pixelix/entity"
)

type challengeRepository struct {
	gormDB *gorm.DB
}

func NewChallengeRepository(gormDB *gorm.DB) *challengeRepository {
	return &challengeRepository{gormDB: gormDB}
}

var _ entity.ChallengeRepository = (*challengeRepository)(nil)

func (cr challengeRepository) CreateChallenge(c context.Context, challenge *entity.Challenge) (*entity.Challenge, error) {
	//TODO implement me
	panic("implement me")
}
