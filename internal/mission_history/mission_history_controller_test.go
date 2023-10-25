package mission_history

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/logger"
	"testing"
)

type controllerTestSuite struct {
	router                   *gin.Engine
	log                      logger.Logger
	missionHistoryService    *mocks.MissionHistoryService
	missionHistoryController entity.MissionHistoryController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.missionHistoryService = mocks.NewMissionHistoryService(t)
	ts.missionHistoryController = NewMissionHistoryController(ts.missionHistoryService, ts.log)
	RegisterRoutes(ts.router, ts.missionHistoryController)

	return ts
}

func Test_missionHistoryController_CreateMissionHistory(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		mock   func()
		status int
	}{
		{
			name: "PASS challenge 생성",
			mock: func() {
				ts.missionHistoryService.EXPECT().
					CreateMissionHistory(mock.Anything, entity.CreateMissionHistoryRequest{}).
					Return(&entity.CreateMissionHistoryResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/mission-histories", nil)

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.missionHistoryService.AssertExpectations(t)
		})
	}
}

func Test_missionHistoryController_ListMissionHistories(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew().String()

	tests := []struct {
		name   string
		mock   func()
		uri    func() string
		query  func() string
		status int
	}{
		{
			name: "PASS challenge 히스토리 조회",
			mock: func() {
				ts.missionHistoryService.EXPECT().
					ListMultiModeMissionHistories(mock.Anything, entity.ListMultiModeMissionHistoriesRequest{
						UserID: testUserID,
						Date:   "2023-10-10",
					}).Return(&entity.ListMultiModeMissionHistoriesResponse{}, nil).Once()
			},
			uri: func() string {
				path, _ := url.JoinPath("/challenge-histories/multi", testUserID)
				return path
			},
			query: func() string {
				params := url.Values{}
				params.Add("date", "2023-10-10")
				return params.Encode()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodGet, tt.uri(), nil)
			req.URL.RawQuery = tt.query()

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.missionHistoryService.AssertExpectations(t)
		})
	}
}
