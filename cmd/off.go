package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var offCmd = &cobra.Command{
	Use:           "off",
	Aliases:       []string{"0"},
	Short:         "fusion reactor control plane",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Powering off: Light %d ...\n", viper.GetInt(flagHueLight))
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		return controller.PowerOff(viper.GetInt(flagHueLight))
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
