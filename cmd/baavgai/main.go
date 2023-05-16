package main

import (
	"github.com/rs/zerolog/log"

	application "baavgai/internal/app"
)

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		log.Fatal().Msg("Error: incorrect application run")
	}
}
