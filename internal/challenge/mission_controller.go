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
	missions := e.Group("/missions")
	{
		missions.POST("", controller.CreateMission)
		missions.GET("/user/:userID", controller.ListMissions)
		missions.GET("/:missionID", controller.GetMission)
		missions.PATCH("", controller.PatchMission)
	}
}

type missionController struct {
	logger  logger.Logger
	service entity.ChallengeService
}

func NewMissionController(service entity.ChallengeService, logger logger.Logger) *missionController {
	return &missionController{
		logger:  logger,
		service: service,
	}
}

var _ entity.ChallengeController = (*missionController)(nil)

func (tc *missionController) CreateMission(c *gin.Context) {
	var req entity.CreateMissionRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.CreateMission(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *missionController) GetMission(c *gin.Context) {
	var req entity.GetMissionRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.GetMission(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *missionController) ListMissions(c *gin.Context) {
	var req entity.ListMissionsRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.ListMissions(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (tc *missionController) PatchMission(c *gin.Context) {
	var req entity.PatchMissionRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := tc.service.PatchMission(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
