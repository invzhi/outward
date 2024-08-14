package outward

import (
	"context"
	"net/http"
	"syscall"

	"github.com/oklog/run"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/api"
	"github.com/invzhi/outward/proto/outward/v1/v1connect"
)

func httpServer(appctx *config.AppContext) (*http.Server, error) {
	mux := http.NewServeMux()

	mux.Handle(v1connect.NewUserServiceHandler(api.NewUserService(appctx)))
	mux.Handle(v1connect.NewWorkspaceServiceHandler(api.NewWorkspaceService(appctx)))

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
