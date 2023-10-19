package single_mode

import (
	"context"
	"github.com/stretchr/testify/mock"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type serviceTestSuite struct {
	missionRepo            *mocks.MissionRepository
	missionParticipantRepo *mocks.MissionParticipantRepository
	missionHistoryRepo     *mocks.MissionHistoryRepository
	service                entity.SingleModeService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.missionRepo = mocks.NewMissionRepository(t)
	ts.missionParticipantRepo = mocks.NewMissionParticipantRepository(t)
	ts.missionHistoryRepo = mocks.NewMissionHistoryRepository(t)
	ts.service = NewSingleModeService(ts.missionRepo, ts.missionParticipantRepo, ts.missionHistoryRepo)

	return ts
}

func Test_singleModeService_CreateSingleModeMissionHistories(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	ts := initServiceTestSuite(t)
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "PASS single mode mission history 생성",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				// 생성 해야할 미션을 찾는다?
				ts.missionRepo.EXPECT().ListActiveSingleMissionIDs(mock.Anything).
					Return([]entity.Mission{}, nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if err := ts.service.CreateSingleModeMissionHistories(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateSingleModeMissionHistories() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
