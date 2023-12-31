package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type logger struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

var _ Logger = (*logger)(nil)

var Module = fx.Module("logger", fx.Provide(fx.Annotate(NewLogger, fx.As(new(Logger)))))

func NewLogger() *logger {
	newLogger, _ := zap.NewProduction()
	return &logger{
		logger:      newLogger,
		sugarLogger: newLogger.Sugar(),
	}
}

func (l *logger) Debug(msg string, fields ...interface{}) {
	l.logger.Debug(msg)
}

func (l *logger) Info(msg string, fields ...interface{}) {
	l.logger.Info(msg)
}

func (l *logger) Warn(msg string, fields ...interface{}) {
	l.logger.Warn(msg)
}

func (l *logger) Error(msg string, fields ...interface{}) {
	l.logger.Error(msg)
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatal(msg)
}
