package challenge_history

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
	router     *gin.Engine
	log        logger.Logger
	service    *mocks.ChallengeHistoryService
	controller entity.ChallengeHistoryController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.service = mocks.NewChallengeHistoryService(t)
	ts.controller = NewChallengeHistoryController(ts.service, ts.log)
	RegisterRoutes(ts.router, ts.controller)

	return ts
}

func Test_challengeHistoryController_CreateMissionHistory(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		mock   func()
		status int
	}{
		{
			name: "PASS challenge 생성",
			mock: func() {
				ts.service.EXPECT().
					CreateMissionHistory(mock.Anything, entity.CreateMissionHistoryRequest{}).
					Return(&entity.CreateMissionHistoryResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/challenges/histories", nil)

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.service.AssertExpectations(t)
		})
	}
}

func Test_challengeHistoryController_ListGroupChallengeHistories(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

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
				ts.service.EXPECT().
					ListGroupChallengeHistories(mock.Anything, entity.ListGroupChallengeHistoriesRequest{
						UserID:      testUserID.String(),
						ChallengeID: 1,
						Date:        "2023-01-01",
					}).Return(&entity.ListGroupChallengeHistoriesResponse{}, nil).Once()
			},
			uri: func() string {
				path, _ := url.JoinPath("/challenges/histories", testUserID.String())
				return path
			},
			query: func() string {
				params := url.Values{}
				params.Add("date", "2023-01-01")
				params.Add("challengeId", "1")
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
			ts.service.AssertExpectations(t)
		})
	}
}
