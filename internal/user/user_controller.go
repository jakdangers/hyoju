package user

import (
	"context"
	"cryptoChallenges/dto"
	"cryptoChallenges/entity"
	"cryptoChallenges/pkg/errors"
	"cryptoChallenges/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Routes(e *gin.Engine, controller entity.UserController) {
	e.POST("/users", controller.CreateUser)
	e.GET("/users", controller.ReadUser)
	e.PUT("/users", controller.UpdateUser)
}

type userController struct {
	logger  log.Logger
	service entity.UserService
}

func NewUserController(service entity.UserService, logger log.Logger) *userController {
	return &userController{
		logger:  logger,
		service: service,
	}
}

var _ entity.UserController = (*userController)(nil)

func (uc *userController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	// TODO 유효성 검증
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := uc.service.CreateUser(ctx, req)
	if err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uc *userController) ReadUser(c *gin.Context) {
	var req dto.ReadUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	// TODO 유효성 검증
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := uc.service.ReadUser(ctx, req)
	if err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	// 0 보다작다 이렇게 핸들링할게 거의없다?
	// 코드로
	// API 개발사 ??

	c.JSON(200, res)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	var req dto.UpdateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	if err := req.Valid(); err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := uc.service.UpdateUser(ctx, req)
	if err != nil {
		c.JSON(errors.ToSentinelAPIError(err))
		return
	}

	c.JSON(200, res)
}
