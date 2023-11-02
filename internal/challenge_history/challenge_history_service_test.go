package challenge_history

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
	challengeRepo            *mocks.ChallengeRepository
	challengeParticipantRepo *mocks.ChallengeParticipantRepository
	challengeHistoryRepo     *mocks.ChallengeHistoryRepository
	userRepo                 *mocks.UserRepository
	service                  entity.ChallengeHistoryService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.challengeRepo = mocks.NewChallengeRepository(t)
	ts.challengeParticipantRepo = mocks.NewChallengeParticipantRepository(t)
	ts.challengeHistoryRepo = mocks.NewChallengeHistoryRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewChallengeHistoryService(ts.challengeRepo, ts.challengeParticipantRepo, ts.challengeHistoryRepo, ts.userRepo)

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
			name: "PASS challenge history 생성",
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

func Test_challengeHistoryService_ListGroupChallengeHistories(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.ListGroupChallengeHistoriesRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.ListGroupChallengeHistoriesResponse
		wantErr bool
	}{
		{
			name: "PASS challenge history 조회",
			args: args{
				ctx: context.Background(),
				req: entity.ListGroupChallengeHistoriesRequest{
					UserID:      testUserID.String(),
					ChallengeID: 1,
					Date:        "2023-01-01",
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.challengeRepo.EXPECT().ListMultiChallenges(mock.Anything, entity.ListMultiChallengeParams{
					UserID:        testUserID,
					StartDateTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Hour * 9),
					Type:          entity.ChallengeTypeGroup,
				}).
					Return([]entity.Challenge{
						{
							Model: gorm.Model{
								ID: 1,
							},
							UserID:    testUserID,
							Title:     "test_mission",
							Emoji:     "test_emoji",
							Duration:  "DAILY",
							StartDate: time.Time{},
							EndDate:   time.Time{},
							PlanTime:  time.Time{},
							Alarm:     true,
							WeekDay:   3,
							Type:      entity.ChallengeTypeGroup,
							Status:    entity.ChallengeStatusActivate,
						},
					}, nil).Once()
				ts.challengeHistoryRepo.EXPECT().ListMultiChallengeHistories(mock.Anything, entity.ListGroupChallengeHistoriesParams{
					UserID:       testUserID,
					ChallengeIDs: []uint{1},
				}).
					Return([]entity.ChallengeHistory{
						{
							Model: gorm.Model{
								ID: 1,
							},
							UserID:      testUserID,
							ChallengeID: 1,
							PlanTime:    time.Date(2023, 10, 10, 10, 00, 10, 00, time.UTC),
							FrontImage:  "front_image",
							BackImage:   "back_image",
						},
					}, nil).Once()
			},
			want: &entity.ListGroupChallengeHistoriesResponse{
				ChallengeHistories: []entity.ChallengeHistoryDTO{
					{
						ID:          1,
						UserID:      testUserID.String(),
						ChallengeID: 1,
						PlanTime:    time.Date(2023, 10, 10, 10, 00, 10, 00, time.UTC),
						FrontImage:  "front_image",
						BackImage:   "back_image",
						Title:       "test_mission",
						Emoji:       "test_emoji",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.ListGroupChallengeHistories(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err, err.Error())
			}
			ts.userRepo.AssertExpectations(t)
			ts.challengeHistoryRepo.AssertExpectations(t)
			ts.challengeRepo.AssertExpectations(t)
			ts.challengeParticipantRepo.AssertExpectations(t)
		})
	}
}
