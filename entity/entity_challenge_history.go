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
	ListMultiChallengeHistories(ctx context.Context, params ListMultipleMissionHistoriesParams) ([]ChallengeHistory, error)
}

type ChallengeHistoryService interface {
	CreateMissionHistory(ctx context.Context, req CreateMissionHistoryRequest) (*CreateMissionHistoryResponse, error)
	ListMultiChallengeHistories(ctx context.Context, req ListMultiChallengeHistoriesRequest) (*ListMultiChallengeHistoriesResponse, error)
}

type ChallengeHistoryController interface {
	CreateMissionHistory(c *gin.Context)
	ListMultiModeMissionHistories(c *gin.Context)
}

type ListMultipleMissionHistoriesParams struct {
	UserID       BinaryUUID
	ChallengeIDs []uint
}
