package entity

import (
	"time"
)

type CreateMissionRequest struct {
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
}

type CreateMissionResponse struct {
	MissionID uint `json:"missionID"`
}

type GetMissionRequest struct {
	MissionID uint `json:"missionID" uri:"missionID"`
}

type GetMissionResponse struct {
	Mission MissionDTO `json:"challenge"`
}

type ListMissionsRequest struct {
	UserID string `json:"userID" uri:"userID"`
}

type ListMissionsResponse struct {
	Missions []MissionDTO `json:"missions"`
}

type PatchMissionRequest struct {
	ID        uint       `json:"MissionID"`
	UserID    string     `json:"userID"`
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

type PatchMissionResponse struct {
	MissionDTO
}

type MissionDTO struct {
	ID        uint      `json:"MissionID"`
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

func MissionDTOFrom(mission Challenge) MissionDTO {

	return MissionDTO{
		ID:        mission.ID,
		AuthorID:  mission.UserID.String(),
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
