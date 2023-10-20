package mission_history

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"mission", fx.Provide(
		fx.Annotate(NewMissionHistoryRepository, fx.As(new(entity.MissionHistoryRepository))),
		fx.Annotate(NewMissionHistoryService, fx.As(new(entity.MissionHistoryService))),
		fx.Annotate(NewMissionHistoryController, fx.As(new(entity.MissionHistoryController))),
	),
)
