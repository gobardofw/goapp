package bootstrap

import (
	appns "__anonymous__/__goapp__/internal/app"

	"github.com/gobardofw/translator"
)

func init() {
	// {{if eq .translator "json"}}
	if t, err := translator.NewJSONTranslator("// {{.locale}}", "./config/strings"); err != nil {
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

	appns.ConfigureMessages(app.Translator())
}