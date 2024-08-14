package config

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type pgxLogAdapter struct{}

func (pgxLogAdapter) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	var l zerolog.Level
	switch level {
	case tracelog.LogLevelTrace:
		l = zerolog.TraceLevel
	case tracelog.LogLevelDebug:
		l = zerolog.DebugLevel
	case tracelog.LogLevelInfo:
		l = zerolog.InfoLevel
	case tracelog.LogLevelWarn:
		l = zerolog.WarnLevel
	case tracelog.LogLevelError:
		l = zerolog.ErrorLevel
	default:
		l = zerolog.NoLevel
	}

	zerolog.Ctx(ctx).WithLevel(l).
		Str("module", "pgx").
		Fields(data).
		Msg(msg)
}
