package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Challenge struct {
	gorm.Model
	UserID        BinaryUUID `db:"user_id"`
	Title         string     `db:"title"`
	Emoji         string     `db:"emoji"`
	StartDate     time.Time  `gorm:"type:timestamp"`
	EndDate       time.Time  `gorm:"type:timestamp"`
	PlanTime      time.Time  `gorm:"type:timestamp"`
	Alarm         bool       `db:"alarm"`
	WeekDay       int        `db:"week_day"`
	Type          string     `db:"type"`
	Status        string     `db:"status"`
	ChallengeCode string     `db:"challenge_code"`
}

type ChallengeRepository interface {
	CreateChallenge(c context.Context, challenge *Challenge) (*Challenge, error)
}

type ChallengeService interface {
	CreateChallenge(c context.Context, req CreateChallengeRequest) (*CreateChallengeResponse, error)
}

type ChallengeController interface {
	CreateChallenge(c *gin.Context)
}
