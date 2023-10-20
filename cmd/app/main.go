package main

import (
	"go.uber.org/fx"
	"pixelix/config"
	"pixelix/internal/mission"
	"pixelix/internal/mission_history"
	"pixelix/internal/mission_participant"
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
		mission_participant.Module,
		mission_history.Module,

		// invoke
		fx.Invoke(
			// service Invoke
			user.RegisterRoutes,
			mission.RegisterRoutes,
			mission_history.RegisterRoutes,
			// Infra Invoke
			server.NewHTTPServer,
		),
	).Run()
}
