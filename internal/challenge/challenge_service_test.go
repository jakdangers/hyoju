package challenge

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"k8s.io/utils/pointer"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/helper"
	"testing"
	"time"
)

type serviceTestSuite struct {
	challengeRepo            *mocks.ChallengeRepository
	challengeParticipantRepo *mocks.MissionParticipantRepository
	userRepo                 *mocks.UserRepository
	service                  entity.ChallengeService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.challengeRepo = mocks.NewChallengeRepository(t)
	ts.challengeParticipantRepo = mocks.NewMissionParticipantRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewChallengeService(ts.challengeRepo, ts.challengeParticipantRepo, ts.userRepo)

	return ts
}

func Test_missionService_CreateChallenge(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.CreateChallengeRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.CreateChallengeResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 생성",
			args: args{
				ctx: context.Background(),
				req: entity.CreateChallengeRequest{
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
					Code:        "test_friendCode",
				}, nil).Once()
				ts.challengeRepo.EXPECT().ChallengeFindByCode(mock.Anything, mock.Anything).Return(nil, nil).Once()
				ts.challengeRepo.EXPECT().CreateChallenge(mock.Anything, mock.Anything).Return(&entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    false,
					WeekDay:  3,
					Type:     "SINGLE",
					Status:   entity.ChallengeStatusDeActivate,
				}, nil).Once()
				ts.challengeParticipantRepo.EXPECT().CreateMissionParticipant(mock.Anything, &entity.ChallengeParticipant{
					UserID:      testUserID,
					ChallengeID: 1,
				}).Return(&entity.ChallengeParticipant{
					UserID:      testUserID,
					ChallengeID: 1,
				}, nil).Once()
			},
			want: &entity.CreateChallengeResponse{
				ChallengeID: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.CreateChallenge(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.challengeRepo.AssertExpectations(t)
			ts.userRepo.AssertExpectations(t)
		})
	}
}

func Test_challengeService_ListChallenges(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.ListChallengesRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.ListChallengesResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 리스트 조회 (타입 SINGLE)",
			args: args{
				ctx: context.Background(),
				req: entity.ListChallengesRequest{
					UserID: testUserID.String(),
					Type:   entity.ChallengeTypeSingle,
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.challengeRepo.EXPECT().ListChallenges(mock.Anything, entity.ListChallengesParams{
					UserID: testUserID,
					Type:   entity.ChallengeTypeSingle,
				}).Return([]entity.Challenge{
					{
						Model: gorm.Model{
							ID: 1,
						},
						UserID:   testUserID,
						Title:    "test_mission",
						Emoji:    "test_emoji",
						Duration: entity.ChallengeDurationDaily,
						Alarm:    false,
						WeekDay:  3,
						Type:     entity.ChallengeTypeSingle,
						Status:   entity.ChallengeStatusActivate,
					},
				}, nil).Once()
			},
			want: &entity.ListChallengesResponse{
				Challenges: []entity.ChallengeDTO{
					{
						ID:       1,
						UserID:   testUserID.String(),
						Title:    "test_mission",
						Emoji:    "test_emoji",
						Duration: string(entity.ChallengeDurationDaily),
						Alarm:    false,
						WeekDay:  []string{"MONDAY", "TUESDAY"},
						Type:     string(entity.ChallengeTypeSingle),
						Status:   string(entity.ChallengeStatusActivate),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "PASS 미션 리스트 조회 (타입 MULTI)",
			args: args{
				ctx: context.Background(),
				req: entity.ListChallengesRequest{
					UserID: testUserID.String(),
					Type:   entity.ChallengeTypeMulti,
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.challengeRepo.EXPECT().ListChallenges(mock.Anything, entity.ListChallengesParams{
					UserID: testUserID,
					Type:   entity.ChallengeTypeMulti,
				}).Return([]entity.Challenge{
					{
						Model: gorm.Model{
							ID: 1,
						},
						UserID:   testUserID,
						Title:    "test_mission",
						Emoji:    "test_emoji",
						Duration: entity.ChallengeDurationDaily,
						Alarm:    false,
						WeekDay:  3,
						Type:     entity.ChallengeTypeMulti,
						Status:   entity.ChallengeStatusActivate,
					},
				}, nil).Once()
			},
			want: &entity.ListChallengesResponse{
				Challenges: []entity.ChallengeDTO{
					{
						ID:       1,
						UserID:   testUserID.String(),
						Title:    "test_mission",
						Emoji:    "test_emoji",
						Duration: string(entity.ChallengeDurationDaily),
						Alarm:    false,
						WeekDay:  []string{"MONDAY", "TUESDAY"},
						Type:     string(entity.ChallengeTypeMulti),
						Status:   string(entity.ChallengeStatusActivate),
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.ListChallenges(tt.args.ctx, tt.args.req)
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
		req entity.PatchChallengeRequest
	}

	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.PatchChallengeResponse
		wantErr bool
	}{
		{
			name: "PASS 미션 수정",
			args: args{
				ctx: context.Background(),
				req: entity.PatchChallengeRequest{
					ID:       1,
					UserID:   testUserID.String(),
					Title:    pointer.String("modified_mission"),
					Emoji:    pointer.String("modified_emoji"),
					Duration: helper.EnumToPointer(entity.ChallengeDurationPeriod),
					Alarm:    pointer.Bool(false),
					WeekDay:  []string{"MONDAY", "TUESDAY", "WEDNESDAY"},
					Type:     helper.EnumToPointer(entity.ChallengeTypeSingle),
					Status:   helper.EnumToPointer(entity.ChallengeStatusActivate),
				},
			},
			mock: func() {
				ts.userRepo.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}, nil).Once()
				ts.challengeRepo.EXPECT().GetChallenge(mock.Anything, uint(1)).Return(&entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "original_mission",
					Emoji:    "original_emoji",
					Duration: entity.ChallengeDurationDaily,
					Alarm:    true,
					WeekDay:  3,
					Type:     entity.ChallengeTypeSingle,
					Status:   entity.ChallengeStatusDeActivate,
				}, nil).Once()
				ts.challengeRepo.EXPECT().PatchChallenge(mock.Anything, &entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.ChallengeDurationPeriod,
					Alarm:    false,
					WeekDay:  7,
					Type:     entity.ChallengeTypeSingle,
					Status:   entity.ChallengeStatusActivate,
				}).Return(&entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.ChallengeDurationPeriod,
					Alarm:    false,
					WeekDay:  7,
					Type:     entity.ChallengeTypeSingle,
					Status:   entity.ChallengeStatusActivate,
				}, nil).Once()
			},
			want: &entity.PatchChallengeResponse{
				ChallengeDTO: entity.ChallengeDTO{
					ID:       1,
					UserID:   testUserID.String(),
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: string(entity.ChallengeDurationPeriod),
					Alarm:    false,
					WeekDay:  []string{"MONDAY", "TUESDAY", "WEDNESDAY"},
					Type:     string(entity.ChallengeTypeSingle),
					Status:   string(entity.ChallengeStatusActivate),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.PatchChallenge(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.challengeRepo.AssertExpectations(t)
			ts.userRepo.AssertExpectations(t)
		})
	}
}
