package main

import (
	"go.uber.org/fx"
	"pixelix/config"
	"pixelix/internal/mission"
	"pixelix/internal/user"
	"pixelix/pkg/db"
	"pixelix/pkg/handler"
	"pixelix/pkg/logger"
	"pixelix/pkg/server"
)

func main() {

	fx.New(
		// infra module
		config.Module,
		db.GormModule,
		handler.Module,
		logger.Module,

		// service module
		user.Module,
		mission.Module,

		// invoke
		fx.Invoke(
			// service Invoke
			user.RegisterRoutes,
			mission.RegisterRoutes,
			// Infra Invoke
			server.NewHTTPServer,
		),
	).Run()
}
