package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

const (
	MissionHistoryStatusInit MissionHistoryStatus = "INIT"
	MissionHistoryStatusDone MissionHistoryStatus = "DONE"
)

type MissionHistoryStatus string

type MissionHistory struct {
	gorm.Model
	UserID     BinaryUUID
	MissionID  uint
	Status     MissionHistoryStatus
	PlanTime   time.Time
	FrontImage string
	BackImage  string
}

func (MissionHistory) TableName() string {
	return "mission_histories"
}

type MissionHistoryRepository interface {
	CreateMissionHistory(ctx context.Context, missionHistory *MissionHistory) (*MissionHistory, error)
	ListMultipleModeMissionHistories(ctx context.Context, params ListMultipleMissionHistoriesParams) ([]MissionHistory, error)
}

type MissionHistoryService interface {
	CreateMissionHistory(ctx context.Context, req CreateMissionHistoryRequest) (*CreateMissionHistoryResponse, error)
	ListMultiModeMissionHistories(ctx context.Context, req ListMultiModeMissionHistoriesRequest) (*ListMultiModeMissionHistoriesResponse, error)
}

type MissionHistoryController interface {
	CreateMissionHistory(c *gin.Context)
	ListMultiModeMissionHistories(c *gin.Context)
}

type ListMultipleMissionHistoriesParams struct {
	UserID     BinaryUUID
	MissionIDs []uint
}
