package main

import (
	"fmt"

	application "baavgai/internal/app"
)

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
