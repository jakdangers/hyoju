package single_mode

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type serviceTestSuite struct {
	missionRepo            *mocks.MissionRepository
	missionParticipantRepo *mocks.MissionParticipantRepository
	missionHistoryRepo     *mocks.MissionHistoryRepository
	userRepo               *mocks.UserRepository
	service                entity.SingleModeService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.missionRepo = mocks.NewMissionRepository(t)
	ts.missionParticipantRepo = mocks.NewMissionParticipantRepository(t)
	ts.missionHistoryRepo = mocks.NewMissionHistoryRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewSingleModeService(ts.missionRepo, ts.missionParticipantRepo, ts.missionHistoryRepo, ts.userRepo)
	return ts
}

func Test_singleModeService_CreateMissionHistories(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	ts := initServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "PASS mission history 생성",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				ts.missionRepo.EXPECT().ListActiveSingleMissionIDs(mock.Anything).Return([]uint{1}, nil).Once()
				ts.missionParticipantRepo.EXPECT().ListMissionParticipants(mock.Anything, uint(1)).
					Return([]entity.MissionParticipant{
						{
							Model: gorm.Model{
								ID: 1,
							},
							UserID:    testUserID,
							MissionID: 1,
						},
					}, nil).Once()
				//ts.missionHistoryRepo.EXPECT()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if err := ts.service.CreateMissionHistories(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateMissionHistories() error = %v, wantErr %v", err, tt.wantErr)
			}
			ts.missionRepo.AssertExpectations(t)
		})
	}
}
