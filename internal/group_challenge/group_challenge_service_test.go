package group_challenge

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type serviceTestSuite struct {
	groupChallengeRepo *mocks.GroupChallengeRepository
	userRepo           *mocks.UserRepository
	service            entity.GroupChallengeService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.userRepo = mocks.NewUserRepository(t)
	ts.groupChallengeRepo = mocks.NewGroupChallengeRepository(t)
	ts.service = NewGroupChallengeService(ts.groupChallengeRepo, ts.userRepo)

	return ts
}

func Test_groupChallengeService_CreateGroupChallenge(t *testing.T) {
	type args struct {
		c   context.Context
		req entity.CreateGroupChallengeRequest
	}

	ts := initServiceTestSuite(t)

	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "PASS 그룹 챌린지 생성",
			mock: func() {
				ts.groupChallengeRepo.EXPECT().CreateGroupChallenge(mock.Anything, &entity.GroupChallenge{
					Title:       "test_group_challenge",
					Description: "test_description",
				}).Return(&entity.GroupChallenge{
					Model: gorm.Model{
						ID: 1,
					},
					Title:       "test_group_challenge",
					Description: "test_description",
				}, nil).Once()
			},
			args: args{
				c: context.Background(),
				req: entity.CreateGroupChallengeRequest{
					Title:       "test_group_challenge",
					Description: "test_description",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := ts.service.CreateGroupChallenge(tt.args.c, tt.args.req)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.groupChallengeRepo.AssertExpectations(t)
		})
	}
}

func Test_groupChallengeService_ListGroupChallenges(t *testing.T) {
	type args struct {
		c   context.Context
		req entity.ListGroupChallengesRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		mock    func()
		args    args
		want    *entity.ListGroupChallengesResponse
		wantErr bool
	}{
		{
			name: "PASS",
			mock: func() {
				ts.groupChallengeRepo.EXPECT().
					ListGroupChallenges(mock.Anything, entity.ListGroupChallengesParams{
						UserID: testUserID,
					}).
					Return(entity.GroupChallenges{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Title:       "test_title",
							Description: "test_description",
						},
					}, nil).Once()
			},
			args: args{
				c: context.Background(),
				req: entity.ListGroupChallengesRequest{
					UserID: testUserID.String(),
				},
			},
			want: &entity.ListGroupChallengesResponse{
				GroupChallenges: []entity.GroupChallengeDto{
					{
						ID:          1,
						Title:       "test_title",
						Description: "test_description",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.ListGroupChallenges(tt.args.c, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
