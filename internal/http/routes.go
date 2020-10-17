package http

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes register web routes
func RegisterRoutes(app *fiber.App) {
	app.Static("/", "./static")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to bardo app")
	})
}
