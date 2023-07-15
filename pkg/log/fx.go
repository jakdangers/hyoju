package log

import (
	"context"
	"cryptoChallenges/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(New, fx.As(new(LoggerWrapper))),
	),
	fx.Invoke(Invoke),
)j

func Invoke(lc fx.Lifecycle, l LoggerWrapper, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Sync()
			return nil
		},
	})
}
