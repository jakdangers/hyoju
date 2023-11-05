package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupChallenge struct {
	gorm.Model
	Title       string
	Description string `db:"description"`
	Emoji       string `db:"emoji"`
}

type GroupChallengeRepository interface {
}

type GroupChallengeService interface {
	CreateGroupChallenge(c context.Context, req CreateGroupChallengeRequest) (CreateGroupChallengeResponse, error)
}

type GroupChallengeController interface {
	CreateGroupChallenge(c *gin.Context)
}
