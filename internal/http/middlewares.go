package http

import (
	"__anonymous__/__goapp__/internal/helpers"
	"time"

	"github.com/gobardofw/cache"
	"github.com/gobardofw/http/middlewares"
	"github.com/gobardofw/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// RegisterMiddlewares register web middlewares
func RegisterMiddlewares(app *fiber.App, c cache.Cache) {
	acLogger := logger.NewLogger(
		"2006-01-02 15:04:05",
		helpers.DateFormatter(),
		logger.NewFileLogger(
			"./storage/access",
			"// {{.name}}",
			"2006-01-02",
			helpers.DateFormatter(),
		),
	)

	app.Use(recover.New())
	app.Use(middlewares.AccessLogger(acLogger))
	app.Use(middlewares.Maintenance(c))
	app.Use(middlewares.RateLimiter("GLOBAL-LIMITER", 60, 1*time.Minute, c))
}
