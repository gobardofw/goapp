package main

import (
	"__anonymous__/__goapp__/internal/bootstrap"
	"__anonymous__/__goapp__/internal/commands"
	"__anonymous__/__goapp__/internal/config"
	"__anonymous__/__goapp__/internal/http"
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
	// {{if eq .database "y"}}
	bootstrap.SetupDatabase()
	// {{end}}
	// {{if eq .web "y"}}
	bootstrap.SetupWeb()
	http.RegisterRoutes()
	bootstrap.App().CLI.AddCommand(commands.ServeCommand)
	// {{end}}
	bootstrap.Run()
}
