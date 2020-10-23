package bootstrap

import (
	"__anonymous__/__goapp__/src/helpers"

	"github.com/gobardofw/cache"
	"github.com/gobardofw/console/maintenance"
	httput "github.com/gobardofw/http"
	"github.com/gobardofw/logger"
	"github.com/gofiber/fiber/v2"
)

// SetupWeb driver
func SetupWeb() {
	conf := fiber.Config{}
	conf.DisableStartupMessage = app.Config().Bool("prod", false)
	conf.ErrorHandler = httput.ErrorLogger(logger.NewLogger(
		"2006-01-02 15:04:05",
		helpers.DateFormatter(),
		logger.NewFileLogger(
			"./logs/error",
			"error",
			"2006-01-02",
			helpers.DateFormatter(),
		),
	), helpers.DateFormatter())
	server := fiber.New(conf)
	_container.Register("--APP-SERVER", server)

	_cli.AddCommand(maintenance.DownCommand(func(driver string) cache.Cache {
		return app.Cache(driver)
	}, "--APP-CACHE"))
	_cli.AddCommand(maintenance.UpCommand(func(driver string) cache.Cache {
		return app.Cache(driver)
	}, "--APP-CACHE"))
}

// Server get web server driver
// leave name empty to resolve default
func (app *AppDriver) Server(names ...string) *fiber.App {
	name := "--APP-SERVER"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(*fiber.App); ok {
			return res
		}
	}
	return nil
}
