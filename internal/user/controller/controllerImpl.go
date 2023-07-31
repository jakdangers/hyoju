package controller

import (
	"cryptoChallenges/internal/user/service"
	"cryptoChallenges/pkg/log"
	"github.com/gin-gonic/gin"
)

type userController struct {
	logger  log.Logger
	service service.UserService
}

func New(service service.UserService, logger log.Logger) *userController {
	return &userController{
		logger:  logger,
		service: service,
	}
}

var _ UserController = (*userController)(nil)

func (uc *userController) CreateUser(ctx *gin.Context) {
	res, _ := uc.service.CreateUser(ctx)
	// logging
	ctx.JSON(200, res)
}

func (uc *userController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
