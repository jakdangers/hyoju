package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

const (
	MissionHistoryStatusInit MissionHistoryStatus = "INIT"
	MissionHistoryStatusFail MissionHistoryStatus = "FAIL"
	MissionHistoryStatusDone MissionHistoryStatus = "DONE"
)

type MissionHistoryStatus string

type MissionHistory struct {
	gorm.Model
	UserID     BinaryUUID
	MissionID  uint
	Status     MissionHistoryStatus
	Date       time.Time
	PlanTime   time.Time
	FrontImage string
	BackImage  string
}

type MissionHistoryRepository interface {
	CreateMissionHistory(ctx context.Context, missionHistory *MissionHistory) (*MissionHistory, error)
	ListMultipleModeMissionHistories(ctx context.Context, params ListMultipleMissionHistoriesParams) ([]MissionHistory, error)
}

type MissionHistoryService interface {
	CreateMissionHistory(ctx context.Context, req CreateMissionHistoryRequest) (*CreateMissionHistoryResponse, error)
	ListMultipleMissionHistories(ctx context.Context, req ListMissionHistoriesRequest) (*ListMissionHistoriesResponse, error)
}

type MissionHistoryController interface {
	CreateMissionHistory(c *gin.Context)
	ListMissionHistories(c *gin.Context)
}

type ListMultipleMissionHistoriesParams struct {
	UserID     BinaryUUID
	MissionIDs []uint
}
