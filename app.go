package outward

import (
	"context"
	"net/http"
	"strings"
	"syscall"

	validate "github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oklog/run"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/api"
	pbv1 "github.com/invzhi/outward/proto/outward/v1"
)

func grpcLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		var l zerolog.Level

		switch level {
		case logging.LevelDebug:
			l = zerolog.DebugLevel
		case logging.LevelInfo:
			l = zerolog.InfoLevel
		case logging.LevelWarn:
			l = zerolog.WarnLevel
		case logging.LevelError:
			l = zerolog.ErrorLevel
		default:
			l = zerolog.NoLevel
		}

		log.WithLevel(l).Ctx(ctx).Fields(fields).Msg(msg)
	})
}

func httpServer(appctx *config.AppContext) (*http.Server, error) {
	// Setup grpc server.
	validator, err := validate.New()
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(grpcLogger()),
			protovalidate.UnaryServerInterceptor(validator),
			recovery.UnaryServerInterceptor(),
		),
	)

	pbv1.RegisterUserServiceServer(server, api.NewUserServer(appctx))
	pbv1.RegisterWorkspaceServiceServer(server, api.NewWorkspaceServer(appctx))

	// Setup grpc gateway.
	ctx := context.Background()
	mux := runtime.NewServeMux()

	conn, err := grpc.NewClient(appctx.Conf.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	if err := pbv1.RegisterUserServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	if err := pbv1.RegisterWorkspaceServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	// Setup http endpoint.
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/grpc") {
			server.ServeHTTP(w, r)
			return
		}

		mux.ServeHTTP(w, r)
	})
	handler := h2c.NewHandler(h, &http2.Server{})

	s := http.Server{Addr: appctx.Conf.Address, Handler: handler}
	s.RegisterOnShutdown(func() { _ = conn.Close() })

	return &s, nil
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
