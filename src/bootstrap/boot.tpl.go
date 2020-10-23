package bootstrap

import (
	"__anonymous__/__goapp__/internal/helpers"
	"os"
	"runtime/debug"
	"time"

	"github.com/gobardofw/cli"
	"github.com/gobardofw/container"
	"github.com/gobardofw/logger"
)

var app *AppDriver
var _container container.Container
var _cli cli.CLI

func init() {
	_container = container.NewContainer()
	_cli = cli.NewCLI("// {{.name}}", "// {{.description}}")
	app = new(AppDriver)
	app.Container = _container
	app.CLI = _cli
}

// App get app instance
func App() *AppDriver {
	return app
}

// Run cli and log panic if exists
func Run() {
	defer func() {
		if r := recover(); r != nil {
			erLogger := logger.NewLogger(
				"2006-01-02 15:04:05",
				helpers.DateFormatter(),
				logger.NewFileLogger(
					"./logs/error",
					"crash",
					"2006-01-02",
					helpers.DateFormatter(),
				),
				os.Stdout,
			)
			erLogger.Divider("=", 100, "APP CRASHED")
			erLogger.Error().Print("%v", r)
			erLogger.Raw("\n\nStacktrace:\n")
			erLogger.Raw(string(debug.Stack()))
			erLogger.Divider("=", 100, helpers.DateFormatter()(time.Now().UTC(), "2006-01-02 15:04:05"))
			erLogger.Raw("\n\n")
		}
	}()
	_cli.Run()
}
