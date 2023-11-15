package group

import (
	"github.com/gin-gonic/gin"
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

func Test_groupController_CreateGroup(t *testing.T) {
	type fields struct {
		groupService entity.GroupService
		logger       logger.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := groupController{
				groupService: tt.fields.groupService,
				logger:       tt.fields.logger,
			}
			g.CreateGroup(tt.args.c)
		})
	}
}
