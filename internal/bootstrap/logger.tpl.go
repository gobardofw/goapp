package bootstrap

import (
	"__anonymous__/__goapp__/internal/helpers"
	"io"
	"os"

	cstorage "github.com/gobardofw/console/storage"
	"github.com/gobardofw/logger"
)

// SetupLogger driver
func SetupLogger() {
	writers := make([]io.Writer, 0)
	writers = append(writers, logger.NewFileLogger("./storage/logs", "// {{.name}}", "2006-01-02", helpers.DateFormatter()))
	if !app.Config().Bool("prod", false) {
		writers = append(writers, os.Stdout)
	}
	if l := logger.NewLogger("2006-01-02 15:04:05", helpers.DateFormatter(), writers...); l != nil {
		_container.Register("--APP-LOGGER", l)
	} else {
		panic("failed to build crypto driver")
	}

	_cli.AddCommand(cstorage.ClearCommand("./storage"))
}
