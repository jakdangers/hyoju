package challenge

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/logger"
	"testing"
)

type controllerTestSuite struct {
	router           *gin.Engine
	log              logger.Logger
	challengeService *mocks.ChallengeService
	controller       entity.ChallengeController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.challengeService = mocks.NewChallengeService(t)
	ts.controller = NewChallengeController(ts.challengeService, ts.log)
	RegisterRoutes(ts.router, ts.controller)

	return ts
}

func Test_challengeController_CreateChallenge(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		mock   func()
		body   func() *bytes.Reader
		status int
	}{
		{
			name: "PASS challenge 생성",
			mock: func() {
				ts.challengeService.EXPECT().CreateChallenge(mock.Anything, entity.CreateChallengeRequest{
					Title:   "test_title",
					Emoji:   "test_emoji",
					Alarm:   false,
					WeekDay: []string{"MONDAY", "TUESDAY"},
				}).Return(&entity.CreateChallengeResponse{}, nil).Once()
			},
			body: func() *bytes.Reader {
				req := entity.CreateChallengeRequest{
					Title:   "test_title",
					Emoji:   "test_emoji",
					Alarm:   false,
					WeekDay: []string{"MONDAY", "TUESDAY"},
				}
				jb, _ := json.Marshal(req)

				return bytes.NewReader(jb)
			},
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/challenges", tt.body())
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.challengeService.AssertExpectations(t)
		})
	}
}
