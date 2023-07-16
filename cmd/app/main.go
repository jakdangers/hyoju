package main

import (
	"cryptoChallenges/config"
	"cryptoChallenges/internal/user"
	"cryptoChallenges/pkg/gorm"
	"cryptoChallenges/pkg/log"
	"cryptoChallenges/pkg/server"
	"cryptoChallenges/pkg/sqlx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Infra modules
		config.Module,
		log.Module,
		gorm.Module,
		sqlx.Module,
		server.Module,
		// service modules
		user.Module,
		// infra Invoke
		fx.Invoke(log.Invoke),
		fx.Invoke(server.Invoke),
		// service Invoke
		fx.Invoke(user.Invoke),
	).Run()
	fx.Options()
}
