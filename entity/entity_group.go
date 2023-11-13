package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
Group
*/
type Group struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	Code        string
}

type GroupRepository interface{}

type GroupService interface {
	CreateGroup(c context.Context, req CreateGroupRequest) error
}

type GroupController interface {
	CreateGroup(c *gin.Context)
}
