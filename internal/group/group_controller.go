package group

import (
	"context"
	"github.com/gin-gonic/gin"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
)

type groupController struct {
	groupService entity.GroupService
	logger       logger.Logger
}

func NewGroupController(groupService entity.GroupService, logger logger.Logger) *groupController {
	return &groupController{groupService: groupService, logger: logger}
}

var _ entity.GroupController = (*groupController)(nil)

func RegisterRoutes(e *gin.Engine, controller entity.GroupController) {
	group := e.Group("/groups")
	{
		group.POST("", controller.CreateGroup)
	}
}

func (g groupController) CreateGroup(c *gin.Context) {
	var req entity.CreateGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := g.groupService.CreateGroup(ctx, req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(200)
}
