package outward

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"syscall"

	"connectrpc.com/connect"
	"github.com/oklog/run"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/api"
	"github.com/invzhi/outward/proto/outward/v1/v1connect"
)

func httpServer(appctx *config.AppContext) (*http.Server, error) {
	opts := []connect.HandlerOption{
		connect.WithInterceptors(api.NewLoggingInterceptor(appctx.Logger)),
		connect.WithRecover(func(_ context.Context, _ connect.Spec, _ http.Header, r any) error {
			return connect.NewError(connect.CodeInternal, errors.New(fmt.Sprint(r)))
		}),
	}

	mux := http.NewServeMux()
	mux.Handle(v1connect.NewUserServiceHandler(api.NewUserService(appctx), opts...))
	mux.Handle(v1connect.NewWorkspaceServiceHandler(api.NewWorkspaceService(appctx), opts...))

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
