package bootstrap

import (
	"github.com/gobardofw/cache"
	"github.com/gobardofw/cli"
	"github.com/gobardofw/config"
	"github.com/gobardofw/container"
	"github.com/gobardofw/crypto"
	"github.com/gobardofw/logger"
	"github.com/gobardofw/translator"
	"github.com/gobardofw/validator"
)

// AppDriver interface
type AppDriver struct {
	Container container.Container
	CLI       cli.CLI
}

// Config get config manager
// leave name empty to resolve default
func (app *AppDriver) Config(names ...string) config.Config {
	name := "--APP-CONFIG"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(config.Config); ok {
			return res
		}
	}
	return nil
}

// Cache get cache manager
// leave name empty to resolve default
func (app *AppDriver) Cache(names ...string) cache.Cache {
	name := "--APP-CACHE"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(cache.Cache); ok {
			return res
		}
	}
	return nil
}

// Crypto get crypto driver
// leave name empty to resolve default
func (app *AppDriver) Crypto(names ...string) crypto.Crypto {
	name := "--APP-CRYPTO"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(crypto.Crypto); ok {
			return res
		}
	}
	return nil
}

// Logger get logger driver
// leave name empty to resolve default
func (app *AppDriver) Logger(names ...string) logger.Logger {
	name := "--APP-LOGGER"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(logger.Logger); ok {
			return res
		}
	}
	return nil
}

// Translator get translator driver
// leave name empty to resolve default
func (app *AppDriver) Translator(names ...string) translator.Translator {
	name := "--APP-TRANSLATOR"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(translator.Translator); ok {
			return res
		}
	}
	return nil
}

// Validator get validator driver
// leave name empty to resolve default
func (app *AppDriver) Validator(names ...string) validator.Validator {
	name := "--APP-VALIDATOR"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(validator.Validator); ok {
			return res
		}
	}
	return nil
}
