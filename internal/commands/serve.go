package commands

import (
	"__anonymous__/__goapp__/internal/bootstrap"
	"__anonymous__/__goapp__/internal/http"

	"github.com/spf13/cobra"
)

func init() {
	bootstrap.App().CLI.AddCommand(
		&cobra.Command{
			Use:   "serve",
			Short: "start web server",
			Run: func(cmd *cobra.Command, args []string) {
				http.RegisterMiddlewares(bootstrap.App().Server())
				http.RegisterRoutes(bootstrap.App().Server())
				bootstrap.App().Server().Listen(bootstrap.App().Config().String("web.port", ":8888"))
			},
		},
	)
}
