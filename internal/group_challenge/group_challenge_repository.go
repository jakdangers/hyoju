package group_challenge

import (
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
