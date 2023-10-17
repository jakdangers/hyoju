package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

const (
	Monday    int = 1 << iota // 1
	Tuesday                   // 4
	Wednesday                 // 8
	Thursday                  // 16
	Friday                    // 32
	Saturday                  // 64
	Sunday                    // 128
)

const (
	Single   = "SINGLE"
	Multiple = "MULTIPLE"
	Active   = "ACTIVE"
	Wait     = "WAIT"
	Daily    = "DAILY"
	Period   = "PERIOD"
)

type Mission struct {
	gorm.Model
	AuthorID  BinaryUUID `db:"author_id"`
	Title     string     `db:"title"`
	Emoji     string     `db:"emoji"`
	Duration  string     `db:"duration"`
	StartDate time.Time  `db:"start_date"`
	EndDate   time.Time  `db:"end_date"`
	PlanTime  time.Time  `db:"plan_time"`
	Alarm     bool       `db:"alarm"`
	WeekDay   int        `db:"week_day"`
	Type      string     `db:"type"`
	Status    string     `db:"status"`
}

type MissionRepository interface {
	CreateMission(ctx context.Context, mission *Mission) (*Mission, error)
	GetMission(ctx context.Context, missionID uint) (*Mission, error)
	ListMissions(ctx context.Context, userID BinaryUUID) ([]Mission, error)
	PatchMission(ctx context.Context, mission *Mission) (*Mission, error)
}

type MissionService interface {
	CreateMission(ctx context.Context, req CreateMissionRequest) (*CreateMissionResponse, error)
	GetMission(ctx context.Context, req GetMissionRequest) (*GetMissionResponse, error)
	ListMissions(ctx context.Context, req ListMissionsRequest) (*ListMissionsResponse, error)
	PatchMission(ctx context.Context, req PatchMissionRequest) (*PatchMissionResponse, error)
}

type MissionController interface {
	CreateMission(c *gin.Context)
	GetMission(c *gin.Context)
	ListMissions(c *gin.Context)
	PatchMission(c *gin.Context)
}

func ConvertDaysOfWeekToInt(daysOfWeek []string) int {
	var result int

	for _, day := range daysOfWeek {
		switch day {
		case "SUNDAY":
			result |= Sunday
		case "MONDAY":
			result |= Monday
		case "TUESDAY":
			result |= Tuesday
		case "WEDNESDAY":
			result |= Wednesday
		case "THURSDAY":
			result |= Thursday
		case "FRIDAY":
			result |= Friday
		case "SATURDAY":
			result |= Saturday
		}
	}

	return result
}

func ConvertIntToDaysOfWeek(days int) []string {
	var selectedDays []string

	dayNames := []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}

	for i := 0; i < len(dayNames); i++ {
		// INT 값에서 해당 비트가 설정되어 있는 경우에만 해당 요일을 선택한 것으로 처리
		if days&(1<<i) != 0 {
			selectedDays = append(selectedDays, dayNames[i])
		}
	}

	return selectedDays
}
