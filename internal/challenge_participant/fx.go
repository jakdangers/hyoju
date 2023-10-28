package challenge_participant

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"challenge_participant", fx.Provide(
		fx.Annotate(NewMissionParticipantRepository, fx.As(new(entity.ChallengeParticipantRepository))),
		//fx.Annotate(NewSer, fx.As(new(entity.ChallengeService))),
		//fx.Annotate(NewMissionController, fx.As(new(entity.ChallengeController))),
	),
)
