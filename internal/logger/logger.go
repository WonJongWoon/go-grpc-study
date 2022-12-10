package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(name string) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build logger for '%s'", name))
	}
	return logger.Named(name)
}

func NewSugar(name string) *zap.SugaredLogger {
	return New(name).Sugar()
}
