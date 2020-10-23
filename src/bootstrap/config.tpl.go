package bootstrap

import "github.com/gobardofw/config"

// SetupConfig driver
func SetupConfig() {
	// {{if eq .config "env"}}
	if c, ok := config.NewEnvConfig("./config/config.env"); ok {
		_container.Register("--APP-CONFIG", c)
	} else {
		panic("failed to build config driver")
	}
	// {{end}}
	// {{if eq .config "json"}}
	if c, ok := config.NewJSONConfig("./config/config.json"); ok {
		_container.Register("--APP-CONFIG", c)
	} else {
		panic("failed to build config driver")
	}
	// {{end}}
	// {{if eq .config "memory"}}
	if c, ok := config.NewMemoryConfig(nil); ok {
		_container.Register("--APP-CONFIG", c)
	} else {
		panic("failed to build config driver")
	}
	// {{end}}
}
