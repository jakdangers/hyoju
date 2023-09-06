package server

import (
	"context"
	"cryptoChallenges/config"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func NewServer(lc fx.Lifecycle, e *gin.Engine, cfg *config.Config) {
	srv := &http.Server{Addr: cfg.HTTP.Port, Handler: e}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			return nil
		},
	})
}
