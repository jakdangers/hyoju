package entity

import "time"

type CreateMissionHistoryRequest struct {
}

type CreateMissionHistoryResponse struct {
}

type ListMissionHistoriesRequest struct {
	UserID string `uri:"userID"`
}

type ListMissionHistoriesResponse struct {
}

type ListGroupChallengeHistoriesRequest struct {
	UserID      string `uri:"userID"`
	ChallengeID uint   `form:"challengeId"`
	Date        string `form:"date"`
}

type ListGroupChallengeHistoriesResponse struct {
	ChallengeHistories []ChallengeHistoryDTO `json:"challengeHistories"`
}

type ChallengeHistoryDTO struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"userId"`
	ChallengeID uint      `json:"challengeId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Emoji       string    `json:"emoji"`
	PlanTime    time.Time `json:"planTime"`
	FrontImage  string    `json:"frontImage"`
	BackImage   string    `json:"backImage"`
}
