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
}

type GroupChallenges []GroupChallenge

type GroupChallengeRepository interface {
	CreateGroupChallenge(c context.Context, groupChallenge *GroupChallenge) (*GroupChallenge, error)
	ListGroupChallenges(c context.Context, req ListGroupChallengesParams) (GroupChallenges, error)
}

type GroupChallengeService interface {
	CreateGroupChallenge(c context.Context, req CreateGroupChallengeRequest) error
	ListGroupChallenges(c context.Context, req ListGroupChallengesRequest) (*ListGroupChallengesResponse, error)
}

type GroupChallengeController interface {
	CreateGroupChallenge(c *gin.Context)
	ListGroupChallenges(c *gin.Context)
}

/*
GroupChallengeRepository
*/

type ListGroupChallengesParams struct {
	UserID BinaryUUID `json:"userID"`
}

/*
GroupChallengeService
*/

/*
GroupChallengeController
*/
