package log

import (
	"cryptoChallenges/config"
	"go.uber.org/zap"
)

type loggerWrapper struct {
	defaultLogger *zap.Logger
	sugarLogger   *zap.SugaredLogger
}

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}

var logger loggerWrapper

var _ LoggerWrapper = (*loggerWrapper)(nil)

func New(cfg *config.Config) *loggerWrapper {
	newLogger, _ := zap.NewProduction()
	logger = loggerWrapper{
		defaultLogger: newLogger,
		sugarLogger:   newLogger.Sugar(),
	}
	return &logger
}

func (l *loggerWrapper) Debug(msg string, fields map[string]interface{}) {
	l.defaultLogger.Debug(msg, zap.Any())
}

func (l *loggerWrapper) Info(msg string, fields map[string]interface{}) {
}

func (l *loggerWrapper) Warn(msg string, fields map[string]interface{}) {
}

func (l *loggerWrapper) Error(msg string, fields map[string]interface{}) {
}

func (l *loggerWrapper) Fatal(msg string, fields map[string]interface{}) {
}
