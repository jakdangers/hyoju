package main

import (
	"cryptoChallenges/config"
	"cryptoChallenges/internal/user"
	"cryptoChallenges/pkg/gorm"
	"cryptoChallenges/pkg/log"
	"cryptoChallenges/pkg/mux"
	"cryptoChallenges/pkg/server"
	"cryptoChallenges/pkg/sqlx"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		// Infra Modules
		config.Module,
		mux.Module,
		log.Module,
		gorm.Module,
		sqlx.Module,
		// Service Modules
		user.Module,
		fx.Invoke(
			// service Invoke
			user.Routes,
			// Infra Invoke
			server.NewServer,
		),
	).Run()
}
