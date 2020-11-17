package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
	"github.com/schnoddelbotz/huego-fe/web"
)

const flagHTTPPort = "http-port"

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:           "serve",
	Aliases:       []string{"s"},
	Short:         "exposes Hue lights control via an ugly web interface",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		return web.Serve(viper.GetString(flagHTTPPort), controller)
	},
}

func init() {
	serveCmd.Flags().StringP(flagHTTPPort, "p", ":1984", "HTTP server TCP [IP]:port")
	_ = viper.BindPFlag(flagHTTPPort, serveCmd.Flags().Lookup(flagHTTPPort))
	rootCmd.AddCommand(serveCmd)
}
