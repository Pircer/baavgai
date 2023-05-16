package app

import (
	"baavgai/internal/config"
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var configFile = flag.String("config", "config.yaml", "config file path")

func (app *Application) LoadConfig() error {
	var err error
	flag.Parse()
	app.Config, err = config.Load(*configFile)
	if err != nil {
		return err
	}
	log.Info().Msg("Configuration loaded successful")
	return nil
}

func (app *Application) SetupLogger() error {
	var logLevel zerolog.Level

	switch app.Config.App.LogLevel {
	case "panic":
		logLevel = zerolog.PanicLevel
	case "fatal":
		logLevel = zerolog.FatalLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "debug":
		logLevel = zerolog.DebugLevel
	case "trace":
		logLevel = zerolog.TraceLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	app.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	return nil
}
