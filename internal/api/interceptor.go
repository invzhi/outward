package api

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/rs/zerolog"
)

func NewLoggingInterceptor(logger zerolog.Logger) connect.Interceptor {
	logger = logger.With().Str("module", "connectrpc").Logger()

	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			startTime := time.Now()

			ctx = logger.WithContext(ctx)
			resp, err := next(ctx, req)

			logger.Info().
				Err(err).
				Dur("latency", time.Since(startTime)).
				Str("procedure", req.Spec().Procedure).
				Str("protocol", req.Peer().Protocol).
				Msg("Unary")

			return resp, err
		}
	})
}
