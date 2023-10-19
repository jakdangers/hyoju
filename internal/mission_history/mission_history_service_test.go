package mission_history

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
	missionRepo            *mocks.MissionRepository
	missionParticipantRepo *mocks.MissionParticipantRepository
	missionHistoryRepo     *mocks.MissionHistoryRepository
	userRepo               *mocks.UserRepository
	service                entity.MissionHistoryService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.missionRepo = mocks.NewMissionRepository(t)
	ts.missionParticipantRepo = mocks.NewMissionParticipantRepository(t)
	ts.missionHistoryRepo = mocks.NewMissionHistoryRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewMissionHistoryService(ts.missionRepo, ts.missionParticipantRepo, ts.missionHistoryRepo, ts.userRepo)

	return ts
}

func Test_missionHistoryService_CreateMissionHistory(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.CreateMissionHistoryRequest
	}

	ts := initServiceTestSuite(t)

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.CreateMissionHistoryResponse
		wantErr bool
	}{
		{
			name: "PASS mission history 생성",
			args: args{
				ctx: context.Background(),
				req: entity.CreateMissionHistoryRequest{},
			},
			mock:    nil,
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.CreateMissionHistory(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err, err.Error())
			}
		})
	}
}

func Test_missionHistoryService_ListMissionHistories(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.ListMissionHistoriesRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.ListMissionHistoriesResponse
		wantErr bool
	}{
		{
			name: "PASS mission history 조회",
			args: args{
				ctx: context.Background(),
				req: entity.ListMissionHistoriesRequest{
					UserID: testUserID.String(),
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.missionRepo.EXPECT().ListMultipleModeMissions(mock.Anything, testUserID).
					Return([]entity.Mission{
						{
							Model: gorm.Model{
								ID: 1,
							},
							AuthorID:  testUserID,
							Title:     "test_mission",
							Emoji:     "test_emoji",
							Duration:  "DAILY",
							StartDate: time.Time{},
							EndDate:   time.Time{},
							PlanTime:  time.Time{},
							Alarm:     true,
							WeekDay:   3,
							Type:      entity.Single,
							Status:    entity.Active,
						},
					}, nil).Once()
				ts.missionHistoryRepo.EXPECT().ListMultipleModeMissionHistories(mock.Anything, entity.ListMultipleMissionHistoriesParams{
					UserID:     testUserID,
					MissionIDs: []uint{1, 2, 3},
				}).
					Return([]entity.MissionHistory{
						{
							Model: gorm.Model{
								ID: 1,
							},
							UserID:     testUserID,
							MissionID:  1,
							Status:     entity.Active,
							Date:       time.Date(2023, 10, 10, 00, 00, 00, 00, time.UTC),
							PlanTime:   time.Date(2023, 10, 10, 10, 00, 10, 00, time.UTC),
							FrontImage: "front_image",
							BackImage:  "back_image",
						},
					}, nil).Once()
			},
			want:    &entity.ListMissionHistoriesResponse{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.ListMultipleMissionHistories(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err, err.Error())
			}
			ts.userRepo.AssertExpectations(t)
			ts.missionHistoryRepo.AssertExpectations(t)
			ts.missionRepo.AssertExpectations(t)
			ts.missionParticipantRepo.AssertExpectations(t)
		})
	}
}
