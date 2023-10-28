package challenge_history

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"challenge_history", fx.Provide(
		fx.Annotate(NewChallengeHistoryRepository, fx.As(new(entity.ChallengeHistoryRepository))),
		fx.Annotate(NewChallengeHistoryService, fx.As(new(entity.ChallengeHistoryService))),
		fx.Annotate(NewChallengeHistoryController, fx.As(new(entity.ChallengeHistoryController))),
	),
)
