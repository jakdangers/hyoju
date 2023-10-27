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
	challenges := e.Group("/challenges")
	{
		challenges.POST("", controller.CreateChallenge)
		challenges.GET("/user/:userId", controller.ListChallenges)
		challenges.GET("/:challengeId", controller.GetChallenge)
		challenges.PATCH("", controller.PatchChallenge)
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

func (tc *challengeController) CreateChallenge(c *gin.Context) {
	var req entity.CreateChallengeRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.CreateChallenge(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *challengeController) GetChallenge(c *gin.Context) {
	var req entity.GetChallengeRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.GetChallenge(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *challengeController) ListChallenges(c *gin.Context) {
	var req entity.ListChallengesRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.ListChallenges(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *challengeController) PatchChallenge(c *gin.Context) {
	var req entity.PatchChallengeRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.PatchChallenge(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
