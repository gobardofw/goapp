package app

import "github.com/gobardofw/config"

// Configure register/override app config
func Configure(config config.Config) {
	config.Set("name", "// {{.name}}")
	config.Set("locale", "// {{.locale}}")
	config.Set("key", "// {{.appKey}}")
}
