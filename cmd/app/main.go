package main

import (
	"cryptoChallenges/config"
	"cryptoChallenges/pkg/db"
	"cryptoChallenges/pkg/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New,
			db.New,
			server.New,
		),
		fx.Invoke(server.Invoke),
	).Run()
}
