package bootstrap

import (
	"github.com/gobardofw/translator"
)

// SetupTranslator driver
func SetupTranslator() {
	// {{if eq .translator "json"}}
	if t, err := translator.NewJSONTranslator("// {{.locale}}", "./config/strings"); err == nil {
		_container.Register("--APP-TRANSLATOR", t)
	} else {
		panic("failed to build translator driver")
	}
	// {{else}}
	if t := translator.NewMemoryTranslator("// {{.locale}}"); t != nil {
		_container.Register("--APP-TRANSLATOR", t)
	} else {
		panic("failed to build translator driver")
	}
	// {{end}}
}
