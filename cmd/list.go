package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		controller.List() // gnah ... this should not print
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
