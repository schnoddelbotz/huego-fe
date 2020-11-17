package cmd

import (
	"fmt"
	"github.com/schnoddelbotz/huego-fe/hue_cmd"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Aliases: []string{"1"},
	Short: "engage rocket launcher",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Powering on: Light %d ...\n", viper.GetInt(flagHueLight))
		return hue_cmd.On(viper.GetString(flagHueIP), viper.GetString(flagHueUser), viper.GetInt(flagHueLight))
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}
