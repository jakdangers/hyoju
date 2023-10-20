package entity

import "time"

type CreateMissionHistoryRequest struct {
}

type CreateMissionHistoryResponse struct {
}

type ListMissionHistoriesRequest struct {
	UserID string `json:"userID" uri:"userID"`
}

type ListMissionHistoriesResponse struct {
}

type ListMultiModeMissionHistoriesRequest struct {
	UserID string `json:"userID" uri:"userID"`
	Date   string `json:"date" form:"date"`
}

type ListMultiModeMissionHistoriesResponse struct {
	MissionHistories []MissionHistoryDTO `json:"missionHistories"`
}

type MissionHistoryDTO struct {
	ID         uint                 `json:"id"`
	UserID     string               `json:"userID"`
	MissionID  uint                 `json:"missionID"`
	Title      string               `json:"title"`
	Emoji      string               `json:"emoji"`
	Status     MissionHistoryStatus `json:"status"`
	PlanTime   time.Time            `json:"planTime"`
	FrontImage string               `json:"frontImage"`
	BackImage  string               `json:"backImage"`
}
