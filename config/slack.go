package config

import "github.com/rs/zerolog/log"

type slackLogAdapter struct{}

func (slackLogAdapter) Output(skip int, s string) error {
	log.Debug().Caller(skip + 1).Msg(s)
	return nil
}
