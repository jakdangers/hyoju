package server

import (
	"context"
	"cryptoChallenges/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func Invoke(lc fx.Lifecycle, g *gin.Engine, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := g.Run(cfg.HTTP.Port); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("error running server: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
