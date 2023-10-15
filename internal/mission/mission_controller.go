package mission

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixelix/dto"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller entity.MissionController) {
	e.POST("/tasks", controller.CreateMission)
}

type taskController struct {
	logger  logger.Logger
	service entity.MissionService
}

func NewTaskController(service entity.MissionService, logger logger.Logger) *taskController {
	return &taskController{
		logger:  logger,
		service: service,
	}
}

var _ entity.MissionController = (*taskController)(nil)

func (tc *taskController) CreateMission(c *gin.Context) {
	var req dto.CreateMissionRequest

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
