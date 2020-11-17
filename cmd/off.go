package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var offCmd = &cobra.Command{
	Use:     "off",
	Aliases: []string{"0"},
	Short:   "fusion reactor control plane",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Powering off: Light %d ...\n", viper.GetInt(flagHueLight))
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		return controller.PowerOff(viper.GetInt(flagHueLight))
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
