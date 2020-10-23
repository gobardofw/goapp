package main

import (
	"__anonymous__/__goapp__/internal/bootstrap"
	"__anonymous__/__goapp__/internal/commands"
	"__anonymous__/__goapp__/internal/config"
	"__anonymous__/__goapp__/internal/helpers"
	"__anonymous__/__goapp__/internal/http"
	"io"
	"os"
	"time"

	"github.com/gobardofw/http/middlewares"
	"github.com/gobardofw/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	bootstrap.SetupConfig()
	config.Configure(bootstrap.App().Config())
	bootstrap.SetupCache()
	bootstrap.SetupCrypto()
	bootstrap.SetupLogger()
	bootstrap.SetupTranslator()
	config.ConfigureMessages(bootstrap.App().Translator())
	bootstrap.SetupValidator()
	// {{if eq .database "y"}}bootstrap.SetupDatabase()// {{end}}
	// {{if eq .web "y"}}
	// Configure web server
	bootstrap.SetupWeb()
	bootstrap.App().Server().Use(recover.New())
	bootstrap.App().Server().Use(middlewares.AccessLogger(createAccessLogger()))
	bootstrap.App().Server().Use(middlewares.Maintenance(bootstrap.App().Cache()))
	bootstrap.App().Server().Use(middlewares.RateLimiter("GLOBAL-LIMITER", 60, 1*time.Minute, bootstrap.App().Cache()))
	http.RegisterRoutes(bootstrap.App().Server())
	bootstrap.App().Server().Static("/", "./public")
	bootstrap.App().CLI.AddCommand(commands.ServeCommand)
	// {{end}}

	// Run App
	bootstrap.Run()
}

func createAccessLogger() logger.Logger {
	writers := make([]io.Writer, 1)
	writers[0] = logger.NewFileLogger("./logs/access", "// {{.name}}", "2006-01-02", helpers.DateFormatter())
	if !bootstrap.App().Config().Bool("prod", false) {
		writers = append(writers, os.Stdout)
	}

	return logger.NewLogger("2006-01-02 15:04:05", helpers.DateFormatter(), writers...)
}
