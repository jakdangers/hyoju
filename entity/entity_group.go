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

type GroupRepository interface {
	CreateGroup(c context.Context, group *Group) (*Group, error)
}

type GroupService interface {
	CreateGroup(c context.Context, req CreateGroupRequest) error
}

type GroupController interface {
	CreateGroup(c *gin.Context)
}
