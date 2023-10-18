package entity

import "context"

type SingleModeService interface {
	CreateMissionHistories(ctx context.Context) error
}

type SingleModeController interface {
}
