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

type ChallengeStatus string
type ChallengeType string
type ChallengeDuration string

const (
	ChallengeStatusActivate   ChallengeStatus   = "ACTIVATE"
	ChallengeStatusDeActivate ChallengeStatus   = "DEACTIVATE"
	ChallengeTypeSingle       ChallengeType     = "SINGLE"
	ChallengeTypeGroup        ChallengeType     = "GROUP"
	ChallengeDurationDaily    ChallengeDuration = "DAILY"
	ChallengeDurationPeriod   ChallengeDuration = "PERIOD"
)

type Challenge struct {
	gorm.Model
	UserID      BinaryUUID        `db:"user_id"`
	Title       string            `db:"title"`
	Description string            `json:"description"`
	Emoji       string            `db:"emoji"`
	StartDate   time.Time         `gorm:"type:timestamp"`
	EndDate     time.Time         `gorm:"type:timestamp"`
	PlanTime    time.Time         `gorm:"type:timestamp"`
	Alarm       bool              `db:"alarm"`
	WeekDay     int               `db:"week_day"`
	Duration    ChallengeDuration `db:"duration"`
	Type        ChallengeType     `db:"type"`
	Status      ChallengeStatus   `db:"status"`
	Code        string            `db:"code"`
}

type ChallengeRepository interface {
	CreateChallenge(ctx context.Context, mission *Challenge) (*Challenge, error)
	GetChallenge(ctx context.Context, challengeID uint) (*Challenge, error)
	ListChallenges(ctx context.Context, params ListChallengesParams) ([]Challenge, error)
	PatchChallenge(ctx context.Context, challenge *Challenge) (*Challenge, error)
	ListMultiChallenges(ctx context.Context, params ListMultiChallengeParams) ([]Challenge, error)
	ChallengeFindByCode(ctx context.Context, code string) (*Challenge, error)
}

type ChallengeService interface {
	CreateChallenge(ctx context.Context, req CreateChallengeRequest) (*CreateChallengeResponse, error)
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

type ListChallengesParams struct {
	UserID BinaryUUID
	Type   ChallengeType
}

type ListMultiChallengeParams struct {
	UserID BinaryUUID
	Date   time.Time
	Type   ChallengeType
}
