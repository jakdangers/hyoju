package server

import (
	"context"
	"cryptoChallenges/config"
	"cryptoChallenges/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(fx.Provide(New))

func Invoke(lc fx.Lifecycle, g *gin.Engine, l log.Logger, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := g.Run(cfg.HTTP.Port); !errors.Is(err, http.ErrServerClosed) {
				l.Fatal(fmt.Sprintf("error running server: %v\n", err))
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
