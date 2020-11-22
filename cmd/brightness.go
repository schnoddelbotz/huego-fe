package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/huecontroller"
)

var brightnessCmd = &cobra.Command{
	Use:           "brightness",
	Aliases:       []string{"b"},
	Short:         "control gravity",
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		brightness, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Setting brightness of light #%d to %d\n",
			viper.GetInt(flagHueLight), uint8(brightness))
		controller := huecontroller.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		if !controller.IsLoggedIn() {
			return errors.New("missing login data; provide as args/env (see -h) or run huego-fe login")
		}
		return controller.SetBrightness(viper.GetInt(flagHueLight), uint8(brightness))
	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)
}
