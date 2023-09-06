package user

import (
	"cryptoChallenges/entity"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewUserRepository, fx.As(new(entity.UserRepository))),
		fx.Annotate(NewUserService, fx.As(new(entity.UserService))),
		fx.Annotate(NewUserController, fx.As(new(entity.UserController))),
	),
)
