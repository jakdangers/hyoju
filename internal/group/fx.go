package group

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"user", fx.Provide(
		fx.Annotate(NewGroupRepository, fx.As(new(entity.GroupRepository))),
		fx.Annotate(NewGroupService, fx.As(new(entity.GroupService))),
		fx.Annotate(NewGroupController, fx.As(new(entity.GroupController))),
	),
)
