package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

var offCmd = &cobra.Command{
	Use:           "off",
	Aliases:       []string{"0"},
	Short:         "fusion reactor control plane",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		if viper.GetBool(flagSingle) {
			fmt.Printf("Powering off: Light %d ...\n", viper.GetInt(flagHueLight))
			return controller.PowerOff(viper.GetInt(flagHueLight))
		}
		fmt.Printf("Powering off: Group %d ...\n", viper.GetInt(flagHueGroup))
		return controller.GroupPowerOff(viper.GetInt(flagHueGroup))
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
