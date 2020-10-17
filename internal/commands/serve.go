package commands

import (
	"github.com/gobardofw/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// ServeCommand start web server
func ServeCommand(conf config.Config, server *fiber.App) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "start web server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Listen(conf.String("web.port", ":8888"))
		},
	}
}
