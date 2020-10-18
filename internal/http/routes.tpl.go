package http

import (
	"__anonymous__/__goapp__/internal/bootstrap"
	"__anonymous__/__goapp__/internal/helpers"
	"time"

	"github.com/gobardofw/http/middlewares"
	"github.com/gobardofw/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// RegisterRoutes register web routes
func RegisterRoutes() {
	app := bootstrap.App().Server()
	// middlewares
	app.Use(recover.New())
	app.Use(middlewares.AccessLogger(createAccessLogger()))
	app.Use(middlewares.Maintenance(bootstrap.App().Cache()))
	app.Use(middlewares.RateLimiter("GLOBAL-LIMITER", 60, 1*time.Minute, bootstrap.App().Cache()))

	// Routes
	app.Static("/", "./static")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to bardo app")
	})
}

func createAccessLogger() logger.Logger {
	return logger.NewLogger(
		"2006-01-02 15:04:05",
		helpers.DateFormatter(),
		logger.NewFileLogger(
			"./storage/access",
			"// {{.name}}",
			"2006-01-02",
			helpers.DateFormatter(),
		),
	)
}
