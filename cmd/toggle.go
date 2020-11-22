package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

var toggleCmd = &cobra.Command{
	Use:           "toggle",
	Aliases:       []string{"t"},
	Short:         "toggle toggles",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: should accept numeric lamp ID from List() ... or look up by name! Order changes...
		fmt.Printf("Toggling light %d ...\n", viper.GetInt(flagHueLight))
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		lightID := viper.GetInt(flagHueLight)
		l, err := controller.LightByID(lightID)
		if err != nil {
			return err
		}
		if l.State.On {
			return controller.PowerOff(lightID)
		}
		return controller.PowerOn(lightID)
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
