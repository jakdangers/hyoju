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

type ListMultiChallengeHistoriesRequest struct {
	UserID string        `json:"userID" uri:"userID"`
	Date   string        `json:"date" form:"date"`
	Type   ChallengeType `json:"type" form:"type"`
}

type ListMultiChallengeHistoriesResponse struct {
	ChallengeHistories []ChallengeHistoryDTO `json:"challengeHistories"`
}

type ChallengeHistoryDTO struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"userID"`
	ChallengeID uint      `json:"challengeID"`
	Title       string    `json:"title"`
	Emoji       string    `json:"emoji"`
	PlanTime    time.Time `json:"planTime"`
	FrontImage  string    `json:"frontImage"`
	BackImage   string    `json:"backImage"`
}
