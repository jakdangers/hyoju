package challenge_history

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller entity.ChallengeHistoryController) {
	histories := e.Group("challenges/histories")
	{
		histories.POST("", controller.CreateMissionHistory)
		histories.GET("/:userID", controller.ListMultiModeMissionHistories)
	}
}

type challengeHistoryController struct {
	logger  logger.Logger
	service entity.ChallengeHistoryService
}

func NewChallengeHistoryController(service entity.ChallengeHistoryService, logger logger.Logger) *challengeHistoryController {
	return &challengeHistoryController{
		logger:  logger,
		service: service,
	}
}

var _ entity.ChallengeHistoryController = (*challengeHistoryController)(nil)

func (m challengeHistoryController) CreateMissionHistory(c *gin.Context) {
	var req entity.CreateMissionHistoryRequest

	res, err := m.service.CreateMissionHistory(c.Request.Context(), req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (m challengeHistoryController) ListMultiModeMissionHistories(c *gin.Context) {
	var req entity.ListMultiChallengeHistoriesRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := m.service.ListMultiChallengeHistories(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
