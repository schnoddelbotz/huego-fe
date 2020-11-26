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
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}

		if viper.GetBool(flagSingle) {
			lightID := viper.GetInt(flagHueLight)
			fmt.Printf("Toggling light %d ...\n", lightID)
			l, err := controller.LightByID(lightID)
			if err != nil {
				return err
			}
			if l.State.On {
				return controller.PowerOff(lightID)
			}
			return controller.PowerOn(lightID)
		}

		groupID := viper.GetInt(flagHueGroup)
		fmt.Printf("Toggling group %d ...\n", groupID)
		l, err := controller.GroupByID(groupID)
		if err != nil {
			return err
		}
		if l.State.On {
			return controller.GroupPowerOff(groupID)
		}
		return controller.GroupPowerOn(groupID)
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
