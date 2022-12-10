package api

import (
	"context"
	"fmt"
	"go-grpc-study/internal/config"
	"go-grpc-study/internal/logger"
	"go-grpc-study/internal/todo"
	pb "go-grpc-study/protos/v1/todo"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

var _logger = logger.NewSugar("api")

type API struct {
	server *grpc.Server
}

type Param struct {
	fx.In
	AppConfig  *config.AppConfig
	TodoServer *todo.Server
}

func New(lc fx.Lifecycle, param Param) *API {
	if param.AppConfig.Server.Port == 0 {
		_logger.Panicf("'server.port' in profile can not be empty or zero")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", param.AppConfig.Server.Port))
	if err != nil {
		_logger.Panicf("failed to open grpc server - port '%d'", param.AppConfig.Server.Port)
	}

	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server, param.TodoServer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Serve(listener); err != nil {
					_logger.Panic("failed to serve grpc service")
				}
			}()
			_logger.Infof("running grpc server on port '%d'", param.AppConfig.Server.Port)
			return nil
		},
	})

	return &API{server: server}
}

func invoke(_ *API) {

}
