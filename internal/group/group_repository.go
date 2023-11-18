package group

import (
	"context"
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

func (g groupRepository) CreateGroup(c context.Context, group *entity.Group) (*entity.Group, error) {
	//TODO implement me
	panic("implement me")
}
