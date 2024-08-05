package config

import (
	"context"
	"fmt"

	zerolog "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"

	"github.com/invzhi/outward/internal/sqlc"
)

type Config struct {
	// Version  conf.Version
	Debug       bool   `conf:"default:false,env:DEBUG"`
	Address     string `conf:"default::8080,env:ADDRESS"`
	DatabaseURL string `conf:"required,env:DATABASE_URL"`
	Slack       struct {
		BotToken      string `conf:"required,env:SLACK_BOT_TOKEN"`
		SigningSecret string `conf:"required,env:SLACK_SIGNING_SECRET"`
	}
}

type AppContext struct {
	Conf    Config
	Pool    *pgxpool.Pool
	Queries *sqlc.Queries
	Slack   *slack.Client
}

func NewAppContext(conf Config) (*AppContext, error) {
	config, err := pgxpool.ParseConfig(conf.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config: %w", err)
	}

	level := tracelog.LogLevelInfo
	if conf.Debug {
		level = tracelog.LogLevelTrace
	}
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   zerolog.NewLogger(log.Logger),
		LogLevel: level,
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	appctx := AppContext{
		Conf:    conf,
		Pool:    pool,
		Queries: sqlc.New(pool),
		Slack: slack.New(conf.Slack.BotToken,
			slack.OptionDebug(conf.Debug),
			slack.OptionLog(slackLogAdapter{}),
		),
	}

	return &appctx, nil
}

type TxFunc func(*sqlc.Queries) error

func (a *AppContext) DoTx(ctx context.Context, f TxFunc) error {
	tx, err := a.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	// nolint: errcheck
	defer tx.Rollback(ctx)

	err = f(a.Queries.WithTx(tx))
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
