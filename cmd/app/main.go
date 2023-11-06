package main

import (
	"go.uber.org/fx"
	"pixelix/config"
	"pixelix/internal/challenge"
	"pixelix/internal/challenge_history"
	"pixelix/internal/challenge_participant"
	"pixelix/internal/group_challenge"
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
		challenge.Module,
		challenge_participant.Module,
		challenge_history.Module,
		group_challenge.Module,

		// invoke
		fx.Invoke(
			// service Invoke
			user.RegisterRoutes,
			challenge.RegisterRoutes,
			challenge_history.RegisterRoutes,
			group_challenge.RegisterRoutes,
			// Infra Invoke
			server.NewHTTPServer,
		),
	).Run()
}
