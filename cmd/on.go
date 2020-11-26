package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

var onCmd = &cobra.Command{
	Use:           "on",
	Aliases:       []string{"1"},
	Short:         "engage rocket launcher",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		if viper.GetBool(flagSingle) {
			fmt.Printf("Powering on: Light %d ...\n", viper.GetInt(flagHueLight))
			return controller.PowerOn(viper.GetInt(flagHueLight))
		}
		fmt.Printf("Powering on: Group %d ...\n", viper.GetInt(flagHueGroup))
		return controller.GroupPowerOn(viper.GetInt(flagHueGroup))
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}
