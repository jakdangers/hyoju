package group

import (
	"context"
	"pixelix/entity"
)

type groupService struct {
	groupRepo entity.GroupRepository
	userRepo  entity.UserRepository
}

func NewGroupService(groupRepo entity.GroupRepository, userRepo entity.UserRepository) entity.GroupService {
	return &groupService{groupRepo: groupRepo, userRepo: userRepo}
}

var _ entity.GroupService = (*groupService)(nil)

func (g groupService) CreateGroup(c context.Context, req entity.CreateGroupRequest) error {
	//TODO implement me
	panic("implement me")
}
