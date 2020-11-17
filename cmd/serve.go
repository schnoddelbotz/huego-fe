package cmd

import (
	"github.com/schnoddelbotz/huego-fe/web"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "runs the thing that philps frogot on the Hoe",
	Run: func(cmd *cobra.Command, args []string) {
		web.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
