package entity

type CreateGroupRequest struct {
	Name        string `json:"name"`
	UserID      string `json:"userId"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type CreateGroupResponse struct {
}
