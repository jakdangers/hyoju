package entity

type CreateGroupChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListGroupChallengesRequest struct {
}

type ListGroupChallengesResponse struct {
}
