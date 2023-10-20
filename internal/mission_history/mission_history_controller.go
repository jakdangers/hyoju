package mission_history

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller entity.MissionHistoryController) {
	missionHistories := e.Group("mission-histories")
	{
		missionHistories.POST("", controller.CreateMissionHistory)
		missionHistories.GET("/multi/:userID", controller.ListMultiModeMissionHistories)
	}
}

type missionHistoryController struct {
	logger  logger.Logger
	service entity.MissionHistoryService
}

func NewMissionHistoryController(service entity.MissionHistoryService, logger logger.Logger) *missionHistoryController {
	return &missionHistoryController{
		logger:  logger,
		service: service,
	}
}

var _ entity.MissionHistoryController = (*missionHistoryController)(nil)

func (m missionHistoryController) CreateMissionHistory(c *gin.Context) {
	var req entity.CreateMissionHistoryRequest

	res, err := m.service.CreateMissionHistory(c.Request.Context(), req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (m missionHistoryController) ListMultiModeMissionHistories(c *gin.Context) {
	var req entity.ListMultiModeMissionHistoriesRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := m.service.ListMultiModeMissionHistories(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
