package group

import (
	"github.com/gin-gonic/gin"
	"pixelix/entity"
	"pixelix/pkg/logger"
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
	c.Status(200)
}
