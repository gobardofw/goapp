package bootstrap

import (
	"github.com/gobardofw/validator"
	"github.com/gobardofw/validator/validations"
)

func init() {
	if v := validator.NewValidator(app.Translator()); v != nil {
		// {{if eq .locale "fa"}}
		validations.RegisterExtraValidations(v)
		// {{end}}
		_container.Register("--APP-VALIDATOR", v)
	} else {
		panic("failed to build validator driver")
	}
}
