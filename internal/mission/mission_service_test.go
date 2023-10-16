package mission

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
	"time"
)

type serviceTestSuite struct {
	missionRepo *mocks.MissionRepository
	userRepo    *mocks.UserRepository
	service     entity.MissionService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.missionRepo = mocks.NewMissionRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewMissionService(ts.missionRepo, ts.userRepo)

	return ts
}

func Test_missionService_CreateMission(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.CreateMissionRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    entity.CreateMissionResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 생성",
			args: args{
				ctx: context.Background(),
				req: entity.CreateMissionRequest{
					UserID:   testUserID.String(),
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    true,
					WeekDay:  []string{"MONDAY", "TUESDAY"},
					Type:     "SINGLE",
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					NickName:    "test_nickName",
					Email:       "test_email",
					Provider:    "test_provider",
					FirebaseUID: "test_firegbaseUID",
					FriendCode:  "test_friendCode",
				}, nil).Once()
				ts.missionRepo.EXPECT().CreateMission(mock.Anything, &entity.Mission{
					Model:    gorm.Model{},
					AuthorID: testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    true,
					WeekDay:  3,
					Type:     "SINGLE",
					Status:   entity.Active,
				}).Return(&entity.Mission{
					Model: gorm.Model{
						ID: 1,
					},
					AuthorID: testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    false,
					WeekDay:  3,
					Type:     "SINGLE",
					Status:   entity.Active,
				}, nil).Once()
			},
			want: entity.CreateMissionResponse{
				ID: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.CreateMission(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.missionRepo.AssertExpectations(t)
			ts.userRepo.AssertExpectations(t)
		})
	}
}

func Test_missionService_ListMissions(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.ListMissionsRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    entity.ListMissionsResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 리스트 조회",
			args: args{
				ctx: context.Background(),
				req: entity.ListMissionsRequest{
					UserID: testUserID.String(),
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.missionRepo.EXPECT().ListMissions(mock.Anything, testUserID).Return([]entity.Mission{
					{
						Model: gorm.Model{
							ID: 1,
						},
						AuthorID: testUserID,
						Title:    "test_mission",
						Emoji:    "test_emoji",
						Duration: entity.Daily,
						Alarm:    false,
						WeekDay:  3,
						Type:     entity.Single,
						Status:   entity.Active,
					},
				}, nil).Once()
			},
			want: entity.ListMissionsResponse{
				Missions: []entity.MissionDTO{
					{
						ID:       1,
						AuthorID: testUserID.String(),
						Title:    "test_mission",
						Emoji:    "test_emoji",
						Duration: entity.Daily,
						Alarm:    false,
						WeekDay:  []string{"MONDAY", "TUESDAY"},
						Type:     entity.Single,
						Status:   entity.Active,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.ListMissions(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.missionRepo.AssertExpectations(t)
			ts.userRepo.AssertExpectations(t)
		})
	}
}

func Test_missionService_PatchMission(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.PatchMissionRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    entity.PatchMissionResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 수정",
			args: args{
				ctx: context.Background(),
				req: entity.PatchMissionRequest{
					ID:       1,
					UserID:   testUserID.String(),
					Title:    "modified_mission",
					Emoji:    "modifed_emoji",
					Duration: entity.Daily,
					Alarm:    false,
					WeekDay:  []string{"MONDAY", "TUESDAY"},
					Type:     entity.Single,
					Status:   entity.Wait,
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
			},
			want:    entity.PatchMissionResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.PatchMission(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.missionRepo.AssertExpectations(t)
			ts.userRepo.AssertExpectations(t)
		})
	}
}
