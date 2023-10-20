package single_mode

import (
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
