package gommonfx

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/agilesoftgrowth/gommon/logger"
	"go.uber.org/fx"
)

type ServerParams struct {
	LC      fx.Lifecycle
	Logger  logger.LoggerService
	Router  http.Handler
	AppName string
	Port    string
}

func NewServer(params ServerParams) *http.Server {
	addr := fmt.Sprintf(":%s", params.Port)
	srv := &http.Server{Addr: addr, Handler: params.Router}

	params.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			params.Logger.Info(fmt.Sprintf("Starting application '%s' on '%s'", params.AppName, addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
