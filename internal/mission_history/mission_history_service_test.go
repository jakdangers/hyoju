package mission_history

import (
	"context"
	"github.com/stretchr/testify/assert"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type serviceTestSuite struct {
	missionRepo            *mocks.MissionRepository
	missionParticipantRepo *mocks.MissionParticipantRepository
	userRepo               *mocks.UserRepository
	service                entity.MissionHistoryService
}

func initServiceTestSuite(t *testing.T) serviceTestSuite {
	var ts serviceTestSuite

	ts.missionRepo = mocks.NewMissionRepository(t)
	ts.missionParticipantRepo = mocks.NewMissionParticipantRepository(t)
	ts.userRepo = mocks.NewUserRepository(t)
	ts.service = NewMissionHistoryService(ts.missionRepo, ts.missionParticipantRepo, ts.userRepo)

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
