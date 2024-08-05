package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/invzhi/outward"
	"github.com/invzhi/outward/config"
)

func main() {
	var c config.Config
	help, err := conf.Parse("", &c)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Print(help)
			return
		}
		fmt.Printf("cannot parse configuration: %s\n", err)
		os.Exit(1)
	}

	log.Logger.Level(zerolog.InfoLevel)
	if c.Debug {
		log.Logger = log.Output(zerolog.NewConsoleWriter()).Level(zerolog.TraceLevel)
	}

	appctx, err := config.NewAppContext(c)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create app context")
	}

	app, err := outward.NewApp(appctx)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create app")
	}

	err = app.Run()
	log.Error().Err(err).Msg("app is stopped")
}
