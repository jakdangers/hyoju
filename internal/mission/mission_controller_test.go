package mission

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
	"time"
)

type controllerTestSuite struct {
	router            *gin.Engine
	log               logger.Logger
	missionService    *mocks.MissionService
	missionController entity.MissionController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.missionService = mocks.NewMissionService(t)
	ts.missionController = NewMissionController(ts.missionService, ts.log)
	RegisterRoutes(ts.router, ts.missionController)

	return ts
}

func Test_missionController_CreateTask(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew().String()

	tests := []struct {
		name   string
		body   func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "PASS mission 생성",
			body: func() *bytes.Reader {
				req := entity.CreateMissionRequest{
					UserID:   testUserID,
					Title:    "tet_mission",
					Emoji:    "test_emoji",
					Duration: entity.Daily,
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    true,
					WeekDay:  []string{"SUNDAY", "MONDAY"},
					Type:     entity.Single,
				}
				jb, _ := json.Marshal(req)

				return bytes.NewReader(jb)
			},
			mock: func() {
				ts.missionService.EXPECT().CreateMission(mock.Anything, entity.CreateMissionRequest{
					UserID:   testUserID,
					Title:    "tet_mission",
					Emoji:    "test_emoji",
					Duration: entity.Daily,
					PlanTime: time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:    true,
					WeekDay:  []string{"SUNDAY", "MONDAY"},
					Type:     entity.Single,
				}).
					Return(entity.CreateMissionResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/mission", tt.body())
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.missionService.AssertExpectations(t)
		})
	}
}

func Test_missionController_ListMissions(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew().String()

	tests := []struct {
		name   string
		url    func() string
		mock   func()
		status int
	}{
		{
			name: "PASS 미션 리스트 조회",
			url: func() string {
				path, _ := url.JoinPath("/mission", testUserID)
				return path
			},
			mock: func() {
				ts.missionService.EXPECT().ListMissions(mock.Anything, entity.ListMissionsRequest{
					UserID: testUserID,
				}).Return(entity.ListMissionsResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodGet, tt.url(), nil)

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.missionService.AssertExpectations(t)
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
				req := entity.PatchMissionRequest{
					ID:       1,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: "DAILY",
					Alarm:    false,
					WeekDay:  []string{"MONDAY", "TUESDAY"},
					Type:     entity.Single,
					Status:   entity.Active,
				}
				jb, _ := json.Marshal(req)

				return bytes.NewReader(jb)
			},
			mock: func() {
				ts.missionService.EXPECT().PatchMission(mock.Anything, entity.PatchMissionRequest{
					ID:       1,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: "DAILY",
					Alarm:    false,
					WeekDay:  []string{"MONDAY", "TUESDAY"},
					Type:     entity.Single,
					Status:   entity.Active,
				}).Return(entity.PatchMissionResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPatch, "/mission", tt.body())
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ts.missionService.AssertExpectations(t)
		})
	}
}
