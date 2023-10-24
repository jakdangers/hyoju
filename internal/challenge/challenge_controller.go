package challenge

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller entity.ChallengeController) {
	challenge := e.Group("/challenges")
	{
		challenge.POST("", controller.CreateChallenge)
	}
}

type challengeController struct {
	logger  logger.Logger
	service entity.ChallengeService
}

func NewChallengeController(service entity.ChallengeService, logger logger.Logger) *challengeController {
	return &challengeController{
		logger:  logger,
		service: service,
	}
}

var _ entity.ChallengeController = (*challengeController)(nil)

func (cc challengeController) CreateChallenge(c *gin.Context) {
	var req entity.CreateChallengeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*30)
	defer cancel()

	res, err := cc.service.CreateChallenge(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
