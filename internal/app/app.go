package app

import (
	"go-grpc-study/internal/api"
	"go-grpc-study/internal/config"
	"go-grpc-study/internal/logger"
	"go-grpc-study/internal/todo"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var _logger = logger.New("app")

func New() *fx.App {
	return fx.New(
		todo.Module,
		api.Module,
		config.Module,
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: _logger,
			}
		}),
	)
}
