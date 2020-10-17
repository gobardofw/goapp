package helpers

import "github.com/gobardofw/logger"

// DateFormatter get default app date formatter
func DateFormatter() logger.TimeFormatter {
	// {{if eq .locale "fa"}}
	return logger.JalaliFormatter
	// {{else}}
	return logger.GregorianFormatter
	// {{end}}
}
