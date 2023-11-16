package group

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/logger"
	"testing"
)

type controllerTestSuite struct {
	router     *gin.Engine
	log        logger.Logger
	service    *mocks.GroupService
	controller entity.GroupController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.service = mocks.NewGroupService(t)
	ts.controller = NewGroupController(ts.service, ts.log)
	RegisterRoutes(ts.router, ts.controller)

	return ts
}

func Test_groupController_CreateGroup(t *testing.T) {
	ts := initControllerTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name   string
		mock   func()
		body   func() *bytes.Reader
		status int
	}{
		{
			name: "PASS group 생성",
			mock: func() {
				//ts.service.EXPECT().CreateGroup(mock.Anything, entity.CreateGroupRequest{
				//	Name:        "test_group",
				//	UserID:      testUserID.String(),
				//	Description: "test_description",
				//	Image:       "test_image",
				//}).Return(nil)
			},
			body: func() *bytes.Reader {
				req := entity.CreateGroupRequest{
					Name:        "test_group",
					UserID:      testUserID.String(),
					Description: "test_description",
					Image:       "test_image",
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
			req, _ := http.NewRequest(http.MethodPost, "/groups", tt.body())
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			//ts.service.AssertExpectations(t)
		})
	}
}
