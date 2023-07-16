package log

import (
	"context"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(fx.Annotate(New, fx.As(new(Logger)))))

func Invoke(lc fx.Lifecycle, l Logger) {
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
