package app

import (
	"baavgai/internal/config"
	"baavgai/internal/transport/http"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Application struct {
	Config      config.Config
	Logger      zerolog.Logger
	HTTPService *fiber.App
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
	app.HTTPService = fiber.New()
	router := http.New(app.HTTPService)
	router.RoutesInit()
	app.Logger.Info().Msg("Application initialization successful")
	return nil
}

func (app *Application) Clear() error {
	app.Logger.Info().Msg("Application clearance successful")
	return nil
}
