package group_challenge

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"group_challenge", fx.Provide(
		fx.Annotate(NewGroupChallengeRepository, fx.As(new(entity.GroupChallengeRepository))),
		fx.Annotate(NewGroupChallengeService, fx.As(new(entity.GroupChallengeService))),
		fx.Annotate(NewGroupChallengeController, fx.As(new(entity.GroupChallengeController))),
	),
)
