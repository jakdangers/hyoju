package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"pixelix/config"
)

type NewHTTPServerParams struct {
	fx.In
	Engine *gin.Engine
	Config *config.Config
}

func NewHTTPServer(lc fx.Lifecycle, params NewHTTPServerParams) {
	srv := &http.Server{Addr: params.Config.HTTP.Port, Handler: params.Engine}

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
