package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
	"github.com/schnoddelbotz/huego-fe/web"
)

const flagHTTPPort = "http-port"
const flagOpenBrowser = "open-browser"

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:           "serve",
	Aliases:       []string{"s"},
	Short:         "exposes Hue lights control via an ugly web interface",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if viper.GetBool(flagOpenBrowser) {
			go func() {
				// let's hope we exit/return on error (from web.Serve) before waking up here...
				time.Sleep(100 * time.Millisecond)
				err := browser.OpenURL(fmt.Sprintf("http://localhost%s", viper.GetString(flagHTTPPort)))
				if err != nil {
					log.Printf("Failed to start your web browser: %s", err)
				}
			}()
		}
		return web.Serve(viper.GetString(flagHTTPPort), controller, Version, viper.GetString(flagLightFilter))
	},
}

func init() {
	serveCmd.Flags().StringP(flagHTTPPort, "p", ":1984", "HTTP server TCP [IP]:port")
	serveCmd.Flags().BoolP(flagOpenBrowser, "o", false, "points your browser to huego-fe")
	_ = viper.BindPFlag(flagHTTPPort, serveCmd.Flags().Lookup(flagHTTPPort))
	_ = viper.BindPFlag(flagOpenBrowser, serveCmd.Flags().Lookup(flagOpenBrowser))
	rootCmd.AddCommand(serveCmd)
}
