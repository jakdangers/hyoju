package user

import (
	"go.uber.org/fx"
	"pixelix/entity"
)

var Module = fx.Module(
	"user", fx.Provide(
		fx.Annotate(NewUserRepository, fx.As(new(entity.UserRepository))),
		fx.Annotate(NewUserService, fx.As(new(entity.UserService))),
		fx.Annotate(NewUserController, fx.As(new(entity.UserController))),
	),
)
