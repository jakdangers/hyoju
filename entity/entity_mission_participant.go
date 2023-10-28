package entity

import (
	"context"
	"gorm.io/gorm"
)

type ChallengeParticipant struct {
	gorm.Model
	UserID      BinaryUUID
	ChallengeID uint
}

type ChallengeParticipantRepository interface {
	CreateChallengeParticipant(ctx context.Context, participant *ChallengeParticipant) (*ChallengeParticipant, error)
	ListMissionParticipants(ctx context.Context, missionID uint) ([]ChallengeParticipant, error)
}

type ChallengeParticipantService interface{}

type ChallengeParticipantController interface{}
