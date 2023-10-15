package mission

import (
	"context"
	"github.com/stretchr/testify/assert"
	"pixelix/dto"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type testSuite struct {
	repository *mocks.MissionRepository
	service    entity.MissionService
}

func setupTestSuite(t *testing.T) testSuite {
	var ts testSuite

	ts.repository = mocks.NewMissionRepository(t)
	ts.service = NewMissionService(ts.repository)

	return ts
}

func Test_missionService_CreateMission(t *testing.T) {
	type args struct {
		ctx context.Context
		req dto.CreateMissionRequest
	}

	ts := setupTestSuite(t)

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    dto.CreateMissionResponse
		wantErr bool
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.service.CreateMission(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			ts.repository.AssertExpectations(t)
		})
	}
}
