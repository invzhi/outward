package outward

import (
	"context"
	"net/http"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oklog/run"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/api"
	proto "github.com/invzhi/outward/proto/outward/v1"
)

func httpServer(appctx *config.AppContext) (*http.Server, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	err := proto.RegisterUserServiceHandlerServer(ctx, mux, api.NewUserServer(appctx))
	if err != nil {
		return nil, err
	}
	err = proto.RegisterWorkspaceServiceHandlerServer(ctx, mux, api.NewWorkspaceServer(appctx))
	if err != nil {
		return nil, err
	}

	return &http.Server{Addr: appctx.Conf.Address, Handler: mux}, nil
}

type App struct {
	g run.Group
}

func NewApp(appctx *config.AppContext) (*App, error) {
	server, err := httpServer(appctx)
	if err != nil {
		return nil, err
	}

	var g run.Group

	ctx := context.Background()
	g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))
	g.Add(server.ListenAndServe, func(error) {
		_ = server.Shutdown(ctx)
	})

	return &App{g: g}, nil
}

func (app *App) Run() error {
	return app.g.Run()
}
