package challenge

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"challenge", fx.Provide(
		fx.Annotate(NewChallengeRepository, fx.As(new(entity.ChallengeRepository))),
		fx.Annotate(NewChallengeService, fx.As(new(entity.ChallengeService))),
		fx.Annotate(NewChallengeController, fx.As(new(entity.ChallengeController))),
	),
)
