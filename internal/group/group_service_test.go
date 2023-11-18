package group

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type serviceTestSuite struct {
	groupRepo *mocks.GroupRepository
	userRepo  *mocks.UserRepository
	service   entity.GroupService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.groupRepo = mocks.NewGroupRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewGroupService(ts.groupRepo, ts.userRepo)

	return ts
}

func Test_groupService_CreateGroup(t *testing.T) {
	type fields struct {
		groupRepo entity.GroupRepository
		userRepo  entity.UserRepository
	}
	type args struct {
		c   context.Context
		req entity.CreateGroupRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := groupService{
				groupRepo: tt.fields.groupRepo,
				userRepo:  tt.fields.userRepo,
			}
			tt.wantErr(t, g.CreateGroup(tt.args.c, tt.args.req), fmt.Sprintf("CreateGroup(%v, %v)", tt.args.c, tt.args.req))
		})
	}
}
