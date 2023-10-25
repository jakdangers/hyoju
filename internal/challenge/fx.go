package challenge

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"challenge", fx.Provide(
		fx.Annotate(NewMissionRepository, fx.As(new(entity.ChallengeRepository))),
		fx.Annotate(NewChallengeService, fx.As(new(entity.ChallengeService))),
		fx.Annotate(NewMissionController, fx.As(new(entity.ChallengeController))),
	),
)
