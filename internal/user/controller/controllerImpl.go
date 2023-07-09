package controller

import (
	"cryptoChallenges/internal/user/service"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func New(service service.UserService) *userController {
	return &userController{service: service}
}

func (uc *userController) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
