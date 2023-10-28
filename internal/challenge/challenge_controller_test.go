package challenge

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/utils/pointer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/helper"
	"pixelix/pkg/logger"
	"testing"
	"time"
)

type controllerTestSuite struct {
	router            *gin.Engine
	log               logger.Logger
	challengeService  *mocks.ChallengeService
	missionController entity.ChallengeController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.challengeService = mocks.NewChallengeService(t)
	ts.missionController = NewChallengeController(ts.challengeService, ts.log)
	RegisterRoutes(ts.router, ts.missionController)

	return ts
}

func Test_missionController_CreateChallenge(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew().String()

	tests := []struct {
		name   string
		body   func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "PASS challenge 생성",
			body: func() *bytes.Reader {
				req := entity.CreateChallengeRequest{
					UserID:   testUserID,
					Title:    "tet_mission",
					Emoji:    "test_emoji",
					Duration: entity.ChallengeDurationDaily,
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    true,
					WeekDay:  []string{"SUNDAY", "MONDAY"},
					Type:     entity.ChallengeTypeSingle,
				}
				jb, _ := json.Marshal(req)

				return bytes.NewReader(jb)
			},
			mock: func() {
				ts.challengeService.EXPECT().CreateChallenge(mock.Anything, entity.CreateChallengeRequest{
					UserID:   testUserID,
					Title:    "tet_mission",
					Emoji:    "test_emoji",
					Duration: entity.ChallengeDurationDaily,
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    true,
					WeekDay:  []string{"SUNDAY", "MONDAY"},
					Type:     entity.ChallengeTypeSingle,
				}).
					Return(&entity.CreateChallengeResponse{}, nil).Once()
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

func Test_challengeController_ListChallenges(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew().String()

	tests := []struct {
		name   string
		url    func() string
		query  func() string
		mock   func()
		status int
	}{
		{
			name: "PASS 미션 리스트 조회 (타입 SINGLE)",
			url: func() string {
				path, _ := url.JoinPath("/challenges/user", testUserID)
				return path
			},
			query: func() string {
				params := url.Values{}
				params.Add("type", string(entity.ChallengeTypeSingle))
				return params.Encode()
			},
			mock: func() {
				ts.challengeService.EXPECT().ListChallenges(mock.Anything, entity.ListChallengesRequest{
					UserID: testUserID,
					Type:   entity.ChallengeTypeSingle,
				}).Return(&entity.ListChallengesResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
		{
			name: "PASS 미션 리스트 조회 (타입 SINGLE)",
			url: func() string {
				path, _ := url.JoinPath("/challenges/user", testUserID)
				return path
			},
			query: func() string {
				params := url.Values{}
				params.Add("type", string(entity.ChallengeTypeMulti))
				return params.Encode()
			},
			mock: func() {
				ts.challengeService.EXPECT().ListChallenges(mock.Anything, entity.ListChallengesRequest{
					UserID: testUserID,
					Type:   entity.ChallengeTypeMulti,
				}).Return(&entity.ListChallengesResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodGet, tt.url(), nil)
			req.URL.RawQuery = tt.query()

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.challengeService.AssertExpectations(t)
		})
	}
}

func Test_missionController_PatchMission(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		body   func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "PASS 미션 수정",
			body: func() *bytes.Reader {
				req := entity.PatchChallengeRequest{
					ID:       1,
					Title:    pointer.String("modified_mission"),
					Emoji:    pointer.String("modified_emoji"),
					Duration: helper.EnumToPointer(entity.ChallengeDurationDaily),
					Alarm:    pointer.Bool(false),
					WeekDay:  []string{"MONDAY", "TUESDAY"},
					Type:     helper.EnumToPointer(entity.ChallengeTypeSingle),
					Status:   helper.EnumToPointer(entity.ChallengeStatusActivate),
				}
				jb, _ := json.Marshal(req)

				return bytes.NewReader(jb)
			},
			mock: func() {
				ts.challengeService.EXPECT().PatchChallenge(mock.Anything, entity.PatchChallengeRequest{
					ID:       1,
					Title:    pointer.String("modified_mission"),
					Emoji:    pointer.String("modified_emoji"),
					Duration: helper.EnumToPointer(entity.ChallengeDurationDaily),
					Alarm:    pointer.Bool(false),
					WeekDay:  []string{"MONDAY", "TUESDAY"},
					Type:     helper.EnumToPointer(entity.ChallengeTypeSingle),
					Status:   helper.EnumToPointer(entity.ChallengeStatusActivate),
				}).Return(&entity.PatchChallengeResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPatch, "/challenges", tt.body())
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.challengeService.AssertExpectations(t)
		})
	}
}

func Test_missionController_GetMission(t *testing.T) {
	ts := initControllerTestSuite(t)

	tests := []struct {
		name   string
		uri    func() string
		mock   func()
		status int
	}{
		{
			name: "PASS 미션 조회",
			uri: func() string {
				path, _ := url.JoinPath("/challenges", "1")
				return path
			},
			mock: func() {
				ts.challengeService.EXPECT().GetChallenge(mock.Anything, entity.GetChallengeRequest{
					ChallengeID: 1,
				}).Return(&entity.GetChallengeResponse{}, nil).Once()
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
			ts.challengeService.AssertExpectations(t)
		})
	}
}
