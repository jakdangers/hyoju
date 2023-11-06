package group_challenge

import (
	"context"
	"github.com/gin-gonic/gin"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
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
	var req entity.CreateGroupChallengeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*30)
	defer cancel()

	err := g.service.CreateGroupChallenge(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(200)
}