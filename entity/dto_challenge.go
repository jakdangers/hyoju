package entity

import (
	"time"
)

type CreateChallengeRequest struct {
	UserID    string            `json:"userId"`
	Title     string            `json:"title"`
	Emoji     string            `json:"emoji"`
	StartDate time.Time         `json:"startDate"`
	EndDate   time.Time         `json:"endDate"`
	PlanTime  time.Time         `json:"planTime"`
	Alarm     bool              `json:"alarm"`
	WeekDay   []string          `json:"weekDay"`
	Type      ChallengeType     `json:"type"`
	Duration  ChallengeDuration `json:"duration"`
}

type CreateChallengeResponse struct {
	ChallengeID uint `json:"challengeId"`
}

type GetChallengeRequest struct {
	ChallengeID uint `json:"ChallengeId" uri:"challengeId"`
}

type GetChallengeResponse struct {
	Challenge ChallengeDto `json:"challenge"`
}

type ListChallengesRequest struct {
	UserID string        `json:"userId" uri:"userId"`
	Type   ChallengeType `json:"type" form:"type"`
}

type ListChallengesResponse struct {
	Challenges []ChallengeDto `json:"challenges"`
}

type PatchChallengeRequest struct {
	ID        uint               `json:"id"`
	UserID    string             `json:"userId"`
	Title     *string            `json:"title"`
	Emoji     *string            `json:"emoji"`
	StartDate *time.Time         `json:"startDate"`
	EndDate   *time.Time         `json:"endDate"`
	PlanTime  *time.Time         `json:"planTime"`
	Alarm     *bool              `json:"alarm"`
	WeekDay   []string           `json:"weekDay"`
	Duration  *ChallengeDuration `json:"duration"`
	Type      *ChallengeType     `json:"type"`
	Status    *ChallengeStatus   `json:"status"`
}

type PatchChallengeResponse struct {
	ChallengeDto
}

type ChallengeDto struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Emoji       string    `json:"emoji"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	PlanTime    time.Time `json:"planTime"`
	Alarm       bool      `json:"alarm"`
	WeekDay     []string  `json:"weekDay"`
	Duration    string    `json:"duration"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
}

func ChallengeDtoFrom(challenge Challenge) ChallengeDto {

	return ChallengeDto{
		ID:          challenge.ID,
		UserID:      challenge.UserID.String(),
		Title:       challenge.Title,
		Description: challenge.Description,
		Emoji:       challenge.Emoji,
		Duration:    string(challenge.Duration),
		StartDate:   challenge.StartDate,
		EndDate:     challenge.EndDate,
		PlanTime:    challenge.PlanTime,
		Alarm:       challenge.Alarm,
		WeekDay:     ConvertIntToDaysOfWeek(challenge.WeekDay),
		Type:        string(challenge.Type),
		Status:      string(challenge.Status),
	}
}
