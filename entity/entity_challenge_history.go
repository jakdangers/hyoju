package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ChallengeHistory struct {
	gorm.Model
	UserID      BinaryUUID
	ChallengeID uint
	PlanTime    time.Time
	FrontImage  string
	BackImage   string
}

func (ChallengeHistory) TableName() string {
	return "challenge_histories"
}

type ChallengeHistoryRepository interface {
	CreateChallengeHistory(ctx context.Context, missionHistory *ChallengeHistory) (*ChallengeHistory, error)
	ListGroupChallengeHistories(ctx context.Context, params ListGroupChallengeHistoriesParams) ([]ChallengeHistory, error)
}

type ChallengeHistoryService interface {
	CreateMissionHistory(ctx context.Context, req CreateMissionHistoryRequest) (*CreateMissionHistoryResponse, error)
	ListGroupChallengeHistories(ctx context.Context, req ListGroupChallengeHistoriesRequest) (*ListGroupChallengeHistoriesResponse, error)
}

type ChallengeHistoryController interface {
	CreateMissionHistory(c *gin.Context)
	ListGroupChallengeHistories(c *gin.Context)
}

type ListGroupChallengeHistoriesParams struct {
	ChallengeID   uint
	StartDateTime time.Time
	EndDateTime   time.Time
}
