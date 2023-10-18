package entity

import (
	"context"
	"gorm.io/gorm"
)

type MissionParticipant struct {
	gorm.Model
	UserID    BinaryUUID
	MissionID uint
}

type MissionParticipantRepository interface {
	CreateMissionParticipant(ctx context.Context, participant *MissionParticipant) (*MissionParticipant, error)
	ListMissionParticipants(ctx context.Context, missionID uint) ([]MissionParticipant, error)
}

type MissionParticipantService interface{}

type MissionParticipantController interface{}
