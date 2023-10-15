package dto

import "time"

type CreateMissionRequest struct {
	Title            string    `json:"title"`
	Emoji            string    `json:"emoji"`
	Duration         string    `json:"duration"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	PlanTime         time.Time `json:"plan_time"`
	Alarm            bool      `json:"alarm"`
	Days             []string  `json:"days"`
	ParticipateUsers []string  `json:"participate_users"`
}

type CreateMissionResponse struct {
	ID string `json:"ID"`
}
