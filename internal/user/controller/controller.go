package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(ctx *gin.Context)
}

func Routes(e *gin.Engine, controller UserController) {
	e.GET("/users", controller.GetUsers)
}
