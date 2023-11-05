package group_challenge

import (
	"github.com/gin-gonic/gin"
	"pixelix/entity"
	"pixelix/pkg/logger"
)

func RegisterRoutes(e *gin.Engine, controller entity.GroupChallengeController) {
	challenges := e.Group("/group-challenges")
	{
		challenges.POST("", controller.CreateGroupChallenge)
	}
}

type groupChallengeController struct {
	logger  logger.Logger
	service entity.GroupChallengeService
}

func NewGroupChallengeController(service entity.GroupChallengeService, logger logger.Logger) *groupChallengeController {
	return &groupChallengeController{
		logger:  logger,
		service: service,
	}
}

var _ entity.GroupChallengeController = (*groupChallengeController)(nil)

func (g groupChallengeController) CreateGroupChallenge(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
