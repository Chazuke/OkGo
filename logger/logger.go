package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type (
	OkGoLogger struct {
		env     string
		log     zerolog.Logger
		service string
	}

	OkGoLogOptions struct {
		WithTimestamp bool
		LogLevel      OkGoLogLevel
	}

	OkGoLogLevel string
)

const (
	Debug OkGoLogLevel = "debug"
	Info  OkGoLogLevel = "info"
	Warn  OkGoLogLevel = "warn"
	Error OkGoLogLevel = "error"
)

func DefaultOpts() *OkGoLogOptions {
	return &OkGoLogOptions{
		WithTimestamp: true,
		LogLevel:      Debug,
	}
}

func New(env string, service string, opts OkGoLogOptions) *OkGoLogger {
	var log = zerolog.New(os.Stdout)

	if opts.WithTimestamp {
		log = log.With().Timestamp().Logger()
	}

	return &OkGoLogger{
		env:     env,
		log:     log,
		service: service,
	}
}
