package http

import "github.com/gofiber/fiber/v2"

func GetHealth(ctx *fiber.Ctx) error {
	response := HealthResponse{
		Status: "OK",
	}
	return ctx.JSON(response)
}
