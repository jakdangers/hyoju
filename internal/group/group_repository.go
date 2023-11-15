package group

import (
	"gorm.io/gorm"
	"pixelix/entity"
)

type groupRepository struct {
	gormDB *gorm.DB
}

func NewGroupRepository(gormDB *gorm.DB) entity.GroupRepository {
	return &groupRepository{gormDB: gormDB}
}

var _ entity.GroupRepository = (*groupRepository)(nil)
