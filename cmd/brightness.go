package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var brightnessCmd = &cobra.Command{
	Use:     "brightness",
	Aliases: []string{"b"},
	Short:   "control gravity",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		brightness, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Setting brightness of light #%d to %d\n",
			viper.GetInt(flagHueLight), uint8(brightness))
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		return controller.SetBrightness(viper.GetInt(flagHueLight), uint8(brightness))
	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)
}
