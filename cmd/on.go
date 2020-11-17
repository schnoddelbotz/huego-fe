package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var onCmd = &cobra.Command{
	Use:           "on",
	Aliases:       []string{"1"},
	Short:         "engage rocket launcher",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: should accept numeric lamp ID from List() ... or look up by name! Order changes...
		fmt.Printf("Powering on: Light %d ...\n", viper.GetInt(flagHueLight))
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		return controller.PowerOn(viper.GetInt(flagHueLight))
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}
