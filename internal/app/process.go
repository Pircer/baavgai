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
		serviceAddress := app.Config.App.Host + ":" + app.Config.App.Port
		if err := app.HTTPService.Listen(serviceAddress); err != nil {
			app.Logger.Error().Msg(err.Error())
		}
	}()

	stop := <-signalListener
	app.Logger.Info().Msgf("Recieved signal: %v", stop)

	if err := app.HTTPService.Shutdown(); err != nil {
		return err
	}
	return nil
}
