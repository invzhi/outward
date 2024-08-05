package outward

import (
	"context"
	"net/http"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oklog/run"
	"google.golang.org/grpc"

	"github.com/invzhi/outward/config"
	proto "github.com/invzhi/outward/proto/outward/v1"
)

type issueServer struct {
	proto.UnimplementedIssueServiceServer
}

func (s *issueServer) GetIssue(ctx context.Context, req *proto.GetIssueRequest) (*proto.Issue, error) {
	return &proto.Issue{Name: "Check boyd"}, nil
}

func httpServer(appctx *config.AppContext) (*http.Server, error) {
	svr := grpc.NewServer()

	server := &issueServer{}
	proto.RegisterIssueServiceServer(svr, server)

	mux := runtime.NewServeMux()
	err := proto.RegisterIssueServiceHandlerServer(context.Background(), mux, server)
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
