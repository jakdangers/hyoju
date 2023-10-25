package challenge

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"k8s.io/utils/pointer"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
	"time"
)

type serviceTestSuite struct {
	challengeRepo          *mocks.MissionRepository
	missionParticipantRepo *mocks.MissionParticipantRepository
	userRepo               *mocks.UserRepository
	service                entity.ChallengeService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.challengeRepo = mocks.NewMissionRepository(t)
	ts.missionParticipantRepo = mocks.NewMissionParticipantRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewChallengeService(ts.challengeRepo, ts.missionParticipantRepo, ts.userRepo)

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
		want    *entity.CreateMissionResponse
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
				ts.challengeRepo.EXPECT().CreateMission(mock.Anything, &entity.Challenge{
					Model:    gorm.Model{},
					UserID:   testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					PlanTime: 15*time.Hour + 30*time.Minute,
					Alarm:    true,
					WeekDay:  3,
					Type:     "SINGLE",
					Status:   entity.Active,
				}).Return(&entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					PlanTime: 15*time.Hour + 30*time.Minute,
					Alarm:    false,
					WeekDay:  3,
					Type:     "SINGLE",
					Status:   entity.Active,
				}, nil).Once()
				ts.missionParticipantRepo.EXPECT().CreateMissionParticipant(mock.Anything, &entity.MissionParticipant{
					UserID:    testUserID,
					MissionID: 1,
				}).Return(&entity.MissionParticipant{
					UserID:    testUserID,
					MissionID: 1,
				}, nil).Once()
			},
			want: &entity.CreateMissionResponse{
				MissionID: 1,
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
			ts.challengeRepo.AssertExpectations(t)
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
		want    *entity.ListMissionsResponse
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
				ts.challengeRepo.EXPECT().ListMissions(mock.Anything, testUserID).Return([]entity.Challenge{
					{
						Model: gorm.Model{
							ID: 1,
						},
						UserID:   testUserID,
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
			want: &entity.ListMissionsResponse{
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
			ts.challengeRepo.AssertExpectations(t)
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
		want    *entity.PatchMissionResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 수정",
			args: args{
				ctx: context.Background(),
				req: entity.PatchMissionRequest{
					ID:       1,
					UserID:   testUserID.String(),
					Title:    pointer.String("modified_mission"),
					Emoji:    pointer.String("modified_emoji"),
					Duration: pointer.String(entity.Period),
					Alarm:    pointer.Bool(false),
					WeekDay:  []string{"MONDAY", "TUESDAY", "WEDNESDAY"},
					Type:     pointer.String(entity.Single),
					Status:   pointer.String(entity.Active),
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.challengeRepo.EXPECT().GetMission(mock.Anything, uint(1)).Return(&entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "original_mission",
					Emoji:    "original_emoji",
					Duration: entity.Daily,
					Alarm:    true,
					WeekDay:  3,
					Type:     entity.Single,
					Status:   entity.Wait,
				}, nil).Once()
				ts.challengeRepo.EXPECT().PatchMission(mock.Anything, &entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.Period,
					Alarm:    false,
					WeekDay:  7,
					Type:     entity.Single,
					Status:   entity.Active,
				}).Return(&entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.Period,
					Alarm:    false,
					WeekDay:  7,
					Type:     entity.Single,
					Status:   entity.Active,
				}, nil).Once()
			},
			want: &entity.PatchMissionResponse{
				MissionDTO: entity.MissionDTO{
					ID:       1,
					AuthorID: testUserID.String(),
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.Period,
					Alarm:    false,
					WeekDay:  []string{"MONDAY", "TUESDAY", "WEDNESDAY"},
					Type:     entity.Single,
					Status:   entity.Active,
				},
			},
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
			ts.challengeRepo.AssertExpectations(t)
			ts.userRepo.AssertExpectations(t)
		})
	}
}
