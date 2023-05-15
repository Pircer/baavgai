package main

import "fmt"

func main() {
	fmt.Println("Application start")

	app := app.New()

	err := app.Run()
	if err != nil {
		return err.Error()
	}
}
