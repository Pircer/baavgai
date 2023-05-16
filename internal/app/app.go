package app

import (
	"baavgai/internal/config"
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

var configFile = flag.String("config", "config.yaml", "config file path")

type Application struct {
	Config config.Config
	Logger zerolog.Logger
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run() error {
	err := app.Init()
	if err != nil {
		return err
	}
	err = app.ProcessWithGracefullShutdown()
	if err != nil {
		return err
	}
	err = app.Clear()
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) Init() error {
	err := app.LoadConfig()
	if err != nil {
		return err
	}
	err = app.SetupLogger()
	if err != nil {
		return err
	}
	app.Logger.Info().Msg("Application initialization successful")
	return nil
}

func (app *Application) Clear() error {
	app.Logger.Info().Msg("Application clearance successful")
	return nil
}

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
	app.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	return nil
}

func (app *Application) ProcessWithGracefullShutdown() error {
	signalListener := make(chan os.Signal, 1)
	defer close(signalListener)

	signal.Notify(signalListener,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	stop := <-signalListener
	app.Logger.Info().Msgf("Recieved signal: %v", stop)
	return nil
}
