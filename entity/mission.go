package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"pixelix/dto"
	"time"
)

const (
	Sunday    byte = 1 << iota // 1
	Monday                     // 2
	Tuesday                    // 4
	Wednesday                  // 8
	Thursday                   // 16
	Friday                     // 32
	Saturday                   // 64
)

type Mission struct {
	Base
	Title     string    `db:"title"`
	Emoji     string    `db:"emoji"`
	Duration  string    `db:"duration"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	PlanTime  time.Time `db:"plan_time"`
	Alarm     bool      `db:"alarm"`
	Days      byte      `db:"days"`
}

type MissionRepository interface {
}

type MissionService interface {
	CreateMission(ctx context.Context, req dto.CreateMissionRequest) (dto.CreateMissionResponse, error)
}

type MissionController interface {
	CreateMission(c *gin.Context)
}
