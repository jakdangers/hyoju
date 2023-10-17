package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller entity.UserController) {
	e.GET("/users", controller.ReadUser)
	e.PUT("/users", controller.UpdateUser)
	e.DELETE("/users/:MissionID", controller.DeleteUser)
	e.POST("/users/login", controller.OAuthLoginUser)
}

type userController struct {
	logger  logger.Logger
	service entity.UserService
}

func NewUserController(service entity.UserService, logger logger.Logger) *userController {
	return &userController{
		logger:  logger,
		service: service,
	}
}

var _ entity.UserController = (*userController)(nil)

func (uc *userController) ReadUser(c *gin.Context) {
	var req entity.ReadUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	// TODO 유효성 검증
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := uc.service.ReadUser(ctx, req)
	if err != nil {
		fmt.Println("dddddd")
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	var req entity.UpdateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Valid(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := uc.service.UpdateUser(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(c *gin.Context) {
	var request entity.DeleteUserRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := uc.service.DeleteUser(ctx, request); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(http.StatusOK)
}

func (uc *userController) OAuthLoginUser(c *gin.Context) {
	var req entity.OAuthLoginUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := uc.service.OAuthLoginUser(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
