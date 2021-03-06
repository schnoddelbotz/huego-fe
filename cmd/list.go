package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

var listCmd = &cobra.Command{
	Use:           "list",
	Aliases:       []string{"l"},
	Short:         "A brief description of your command",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		controller.List()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
