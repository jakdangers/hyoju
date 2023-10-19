package entity

import "context"

type SingleModeService interface {
	CreateSingleModeMissionHistories(ctx context.Context) error
}

type SingleModeController interface{}
