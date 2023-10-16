package entity

import (
	"time"
)

type CreateMissionRequest struct {
	UserID    string    `json:"userID"`
	Title     string    `json:"title"`
	Emoji     string    `json:"emoji"`
	Duration  string    `json:"duration"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	PlanTime  time.Time `json:"plan_time"`
	Alarm     bool      `json:"alarm"`
	WeekDay   []string  `json:"weekDay"`
	Type      string    `json:"type"`
}

type CreateMissionResponse struct {
	ID uint `json:"ID"`
}

type ListMissionsRequest struct {
	UserID string `json:"userID" uri:"userID"`
}

type ListMissionsResponse struct {
	Missions []MissionDTO `json:"missions"`
}

type PatchMissionRequest struct {
	ID        uint      `json:"ID"`
	UserID    string    `json:"userID"`
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

type PatchMissionResponse struct{}

type MissionDTO struct {
	ID        uint      `json:"ID"`
	AuthorID  string    `json:"authorID"`
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

func MissionDTOFrom(mission Mission) MissionDTO {
	return MissionDTO{
		ID:        mission.ID,
		AuthorID:  mission.AuthorID.String(),
		Title:     mission.Title,
		Emoji:     mission.Emoji,
		Duration:  mission.Duration,
		StartDate: mission.StartDate,
		EndDate:   mission.EndDate,
		PlanTime:  mission.PlanTime,
		Alarm:     mission.Alarm,
		WeekDay:   ConvertIntToDaysOfWeek(mission.WeekDay),
		Type:      mission.Type,
		Status:    mission.Status,
	}
}
