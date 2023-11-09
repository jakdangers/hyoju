package group_challenge

import (
	"bytes"
	"encoding/json"
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
	service    *mocks.GroupChallengeService
	controller entity.GroupChallengeController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.service = mocks.NewGroupChallengeService(t)
	ts.controller = NewGroupChallengeController(ts.service, ts.log)
	RegisterRoutes(ts.router, ts.controller)

	return ts
}

func Test_groupChallengeController_CreateGroupChallenge(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		mock   func()
		body   func() *bytes.Reader
		status int
	}{
		{
			name: "PASS group challenge 생성",
			mock: func() {
				ts.service.EXPECT().CreateGroupChallenge(mock.Anything, entity.CreateGroupChallengeRequest{
					Title:       "test_group_challenge",
					Description: "test_description",
				}).Return(nil).Once()
			},
			body: func() *bytes.Reader {
				req := entity.CreateGroupChallengeRequest{
					Title:       "test_group_challenge",
					Description: "test_description",
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
			req, _ := http.NewRequest(http.MethodPost, "/group-challenges", tt.body())

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.service.AssertExpectations(t)
		})
	}
}

func Test_groupChallengeController_ListGroupChallenge(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		mock   func()
		uri    func() string
		status int
	}{
		{
			name: "PASS 그룹 챌린지 목록 조회",
			mock: func() {
				ts.service.EXPECT().
					ListGroupChallenges(mock.Anything, entity.ListGroupChallengesRequest{}).
					Return(&entity.ListGroupChallengesResponse{}, nil).Once()
			},
			uri: func() string {
				path, _ := url.JoinPath("/group-challenges", "1")
				return path
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodGet, tt.uri(), nil)

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.service.AssertExpectations(t)
		})
	}
}
