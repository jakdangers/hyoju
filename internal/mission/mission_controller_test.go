package mission

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"pixelix/dto"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/logger"
	"testing"
	"time"
)

type missionControllerTestSuite struct {
	router         *gin.Engine
	log            logger.Logger
	taskService    *mocks.TaskService
	taskController entity.MissionController
}

func setupMissionControllerTestSuite(t *testing.T) missionControllerTestSuite {
	var ms missionControllerTestSuite

	gin.SetMode(gin.TestMode)
	ms.router = gin.Default()
	ms.taskService = mocks.NewTaskService(t)
	ms.taskController = NewTaskController(ms.taskService, ms.log)
	RegisterRoutes(ms.router, ms.taskController)

	return ms
}

func Test_missionController_CreateTask(t *testing.T) {

	ms := setupMissionControllerTestSuite(t)
	testMeID := entity.BinaryUUIDNew().String()

	tests := []struct {
		name   string
		body   func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "PASS mission 생성",
			body: func() *bytes.Reader {
				req := dto.CreateMissionRequest{
					Title:            "tet_mission",
					Emoji:            "test_emoji",
					Duration:         "DAILY",
					StartDate:        time.Time{},
					EndDate:          time.Time{},
					PlanTime:         time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:            true,
					Days:             []string{"SUNDAY", "MONDAY"},
					ParticipateUsers: []string{testMeID},
				}
				jb, _ := json.Marshal(req)

				return bytes.NewReader(jb)
			},
			mock: func() {
				ms.taskService.EXPECT().CreateTask(mock.Anything, dto.CreateMissionRequest{
					Title:            "tet_mission",
					Emoji:            "test_emoji",
					Duration:         "DAILY",
					StartDate:        time.Time{},
					EndDate:          time.Time{},
					PlanTime:         time.Date(2023, time.October, 14, 15, 30, 0, 0, time.UTC),
					Alarm:            true,
					Days:             []string{"SUNDAY", "MONDAY"},
					ParticipateUsers: []string{testMeID},
				}).
					Return(dto.CreateMissionResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/tasks", tt.body())
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ms.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			ms.taskService.AssertExpectations(t)
		})
	}
}
