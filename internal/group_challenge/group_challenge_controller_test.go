package group_challenge

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
	service    *mocks.GroupChallengeService
	controller entity.GroupChallengeController
}

func initControllerTestSuite(t *testing.T) controllerTestSuite {
	var ts controllerTestSuite

	gin.SetMode(gin.TestMode)
	ts.router = gin.Default()
	ts.service = mocks.NewGroupChallengeService(t)
	ts.controller = NewGroupChallengeController(ts.service, ts.log)

	return ts
}
