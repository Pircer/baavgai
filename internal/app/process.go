package app

import (
	"os"
	"os/signal"
	"syscall"
)

func (app *Application) ProcessWithGracefullShutdown() error {
	signalListener := make(chan os.Signal, 1)
	defer close(signalListener)

	signal.Notify(signalListener,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		serviceAddress := ":" + app.Config.App.Port
		app.Logger.Info().Msgf("HTPP service port: %v", serviceAddress)
		if err := app.HTTPService.Listen(serviceAddress); err != nil {
			app.Logger.Error().Msg(err.Error())
		}
	}()

	stop := <-signalListener
	app.Logger.Info().Msgf("Recieved signal: %v", stop)

	return nil
}
