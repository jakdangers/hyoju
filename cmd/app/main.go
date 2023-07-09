package main

import (
	"cryptoChallenges/config"
	"cryptoChallenges/internal/user"
	"cryptoChallenges/pkg/gorm"
	"cryptoChallenges/pkg/server"
	"cryptoChallenges/pkg/sqlx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New,
			gorm.New,
			sqlx.New,
			server.New,
		),
		user.Module,
		fx.Invoke(server.Invoke),
	).Run()
	fx.Options()
}
