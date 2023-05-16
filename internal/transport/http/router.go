package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	App *fiber.App
}

func New(app *fiber.App) *Router {
	return &Router{
		App: app,
	}
}

func (router *Router) RoutesInit() {
	router.App.Get("/health/", GetHealth)
}
