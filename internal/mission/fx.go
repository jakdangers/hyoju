package mission

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"mission", fx.Provide(
		fx.Annotate(NewMissionRepository, fx.As(new(entity.MissionRepository))),
		fx.Annotate(NewMissionService, fx.As(new(entity.MissionService))),
		fx.Annotate(NewMissionController, fx.As(new(entity.MissionController))),
	),
)
