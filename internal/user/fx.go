package user

import (
	"cryptoChallenges/internal/user/controller"
	"cryptoChallenges/internal/user/repository"
	"cryptoChallenges/internal/user/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(repository.New, fx.As(new(repository.UserRepository))),
		fx.Annotate(service.New, fx.As(new(service.UserService))),
		fx.Annotate(controller.New, fx.As(new(controller.UserController))),
	),
)

var Invoke = controller.Routes
