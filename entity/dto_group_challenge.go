package entity

type GroupChallengeDto struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GroupChallengeDtoFrom(groupChallenge GroupChallenge) GroupChallengeDto {
	return GroupChallengeDto{
		ID:          groupChallenge.ID,
		Title:       groupChallenge.Title,
		Description: groupChallenge.Description,
	}
}

/*
GroupChallengeRepository
*/

type CreateGroupChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

/*
GroupChallengeServices
*/

type ListGroupChallengesRequest struct {
	UserID string `json:"userID"`
}

type ListGroupChallengesResponse struct {
	GroupChallenges []GroupChallengeDto `json:"groupChallenges"`
}

/*
GroupChallengeController
*/
