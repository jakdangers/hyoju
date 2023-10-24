package entity

import "time"

type CreateChallengeRequest struct {
	Title     string    `json:"title"`
	Emoji     string    `json:"emoji"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PlanTime  time.Time `json:"planTime"`
	Alarm     bool      `json:"alarm"`
	WeekDay   []string  `json:"weekDay"`
}

type CreateChallengeResponse struct {
}
