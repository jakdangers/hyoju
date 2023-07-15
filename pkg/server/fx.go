package server

import (
	"context"
	"cryptoChallenges/config"
	"cryptoChallenges/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"net/http"
)

func Invoke(lc fx.Lifecycle, g *gin.Engine, l *log.LoggerWrapper, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := g.Run(cfg.HTTP.Port); !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("error running server: %v\n", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
