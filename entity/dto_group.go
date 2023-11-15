package entity

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type CreateGroupResponse struct {
}
