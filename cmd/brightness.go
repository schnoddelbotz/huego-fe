package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hue_cmd"
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
		return hue_cmd.Brightness(viper.GetString(flagHueIP), viper.GetString(flagHueUser),
			viper.GetInt(flagHueLight), uint8(brightness))
	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)
}
