package bootstrap

import (
	"__anonymous__/__goapp__/internal/commands"
	"__anonymous__/__goapp__/internal/helpers"
	"__anonymous__/__goapp__/internal/http"

	"github.com/gobardofw/cache"
	"github.com/gobardofw/console/maintenance"
	httput "github.com/gobardofw/http"
	"github.com/gobardofw/logger"
	"github.com/gofiber/fiber/v2"
)

func init() {
	conf := fiber.Config{}
	conf.DisableStartupMessage = app.Config().Bool("prod", false)
	conf.ErrorHandler = httput.ErrorLogger(logger.NewLogger(
		"2006-01-02 15:04:05",
		helpers.DateFormatter(),
		logger.NewFileLogger(
			"./storage/errors",
			"error",
			"2006-01-02",
			helpers.DateFormatter(),
		),
	), helpers.DateFormatter())
	server := fiber.New(conf)
	_container.Register("--APP-SERVER", server)
	http.RegisterMiddlewares(server, app.Cache())
	http.RegisterRoutes(server)

	_cli.AddCommand(commands.ServeCommand(app.Config(), server))
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
