package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

func Routes(e *gin.Engine, controller UserController) {
	e.POST("/users", controller.CreateUser)
	e.GET("/users", controller.GetUser)
}
