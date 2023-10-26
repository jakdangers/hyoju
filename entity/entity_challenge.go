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
	Single = "SINGLE"
	Multi  = "MULTI"
	Active = "ACTIVE"
	Wait   = "WAIT"
	Daily  = "DAILY"
	Period = "PERIOD"
)

type ChallengeStatus string

const (
	ChallengeStatusActivate   ChallengeStatus = "ACTIVATE"
	ChallengeStatusDeActivate ChallengeStatus = "DEACTIVATE"
)

type Challenge struct {
	gorm.Model
	UserID    BinaryUUID      `db:"user_id"`
	Title     string          `db:"title"`
	Emoji     string          `db:"emoji"`
	Duration  string          `db:"duration"`
	StartDate time.Time       `gorm:"type:timestamp"`
	EndDate   time.Time       `gorm:"type:timestamp"`
	PlanTime  time.Time       `gorm:"type:timestamp"`
	Alarm     bool            `db:"alarm"`
	WeekDay   int             `db:"week_day"`
	Type      string          `db:"type"`
	Status    ChallengeStatus `db:"status"`
}

type ChallengeRepository interface {
	CreateChallenge(ctx context.Context, mission *Challenge) (*Challenge, error)
	GetChallenge(ctx context.Context, missionID uint) (*Challenge, error)
	ListChallenges(ctx context.Context, userID BinaryUUID) ([]Challenge, error)
	PatchChallenge(ctx context.Context, mission *Challenge) (*Challenge, error)
	ListMultiModeMissions(ctx context.Context, params ListMultiModeMissionsParams) ([]Challenge, error)
}

type ChallengeService interface {
	CreateChallenge(ctx context.Context, req CreateChallengeRequest) (*CreateMissionResponse, error)
	GetChallenge(ctx context.Context, req GetChallengeRequest) (*GetChallengeResponse, error)
	ListChallenges(ctx context.Context, req ListChallengesRequest) (*ListChallengesResponse, error)
	PatchChallenge(ctx context.Context, req PatchChallengeRequest) (*PatchChallengeResponse, error)
}

type ChallengeController interface {
	CreateChallenge(c *gin.Context)
	GetChallenge(c *gin.Context)
	ListChallenges(c *gin.Context)
	PatchChallenge(c *gin.Context)
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

type ListMultiModeMissionsParams struct {
	UserID BinaryUUID
	Date   time.Time
}
