package mission_participant

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"missionParticipant", fx.Provide(
		fx.Annotate(NewMissionParticipantRepository, fx.As(new(entity.MissionParticipantRepository))),
		//fx.Annotate(NewSer, fx.As(new(entity.MissionService))),
		//fx.Annotate(NewMissionController, fx.As(new(entity.MissionController))),
	),
)
