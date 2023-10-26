package entity

import (
	"time"
)

type CreateChallengeRequest struct {
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	Emoji     string    `json:"emoji"`
	Duration  string    `json:"duration"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PlanTime  time.Time `json:"planTime"`
	Alarm     bool      `json:"alarm"`
	WeekDay   []string  `json:"weekDay"`
	Type      string    `json:"type"`
}

type CreateMissionResponse struct {
	ChallengeID uint `json:"challengeId"`
}

type GetChallengeRequest struct {
	ChallengeID uint `json:"ChallengeId" uri:"challengeId"`
}

type GetChallengeResponse struct {
	Challenge ChallengeDTO `json:"challenge"`
}

type ListChallengesRequest struct {
	UserID string `json:"userId" uri:"userId"`
}

type ListChallengesResponse struct {
	Challenges []ChallengeDTO `json:"challenges"`
}

type PatchChallengeRequest struct {
	ID        uint       `json:"id"`
	UserID    string     `json:"userId"`
	Title     *string    `json:"title"`
	Emoji     *string    `json:"emoji"`
	Duration  *string    `json:"duration"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
	PlanTime  *time.Time `json:"planTime"`
	Alarm     *bool      `json:"alarm"`
	WeekDay   []string   `json:"weekDay"`
	Type      *string    `json:"type"`
	Status    *string    `json:"status"`
}

type PatchChallengeResponse struct {
	ChallengeDTO
}

type ChallengeDTO struct {
	ID        uint      `json:"id"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	Emoji     string    `json:"emoji"`
	Duration  string    `json:"duration"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PlanTime  time.Time `json:"planTime"`
	Alarm     bool      `json:"alarm"`
	WeekDay   []string  `json:"weekDay"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
}

func ChallengeDTOFrom(challenge Challenge) ChallengeDTO {

	return ChallengeDTO{
		ID:        challenge.ID,
		UserID:    challenge.UserID.String(),
		Title:     challenge.Title,
		Emoji:     challenge.Emoji,
		Duration:  challenge.Duration,
		StartDate: challenge.StartDate,
		EndDate:   challenge.EndDate,
		PlanTime:  challenge.PlanTime,
		Alarm:     challenge.Alarm,
		WeekDay:   ConvertIntToDaysOfWeek(challenge.WeekDay),
		Type:      challenge.Type,
		Status:    string(challenge.Status),
	}
}
