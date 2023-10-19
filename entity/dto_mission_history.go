package entity

type CreateMissionHistoryRequest struct {
}

type CreateMissionHistoryResponse struct {
}

type ListMissionHistoriesRequest struct {
	UserID string `json:"userID" uri:"userID"`
}

type ListMissionHistoriesResponse struct {
}
