package main

import (
	"__anonymous__/__goapp__/internal/bootstrap"
	"__anonymous__/__goapp__/internal/config"
)

func main() {
	config.Configure(bootstrap.App().Config())
	config.ConfigureMessages(bootstrap.App().Translator())
	bootstrap.Run()
}
