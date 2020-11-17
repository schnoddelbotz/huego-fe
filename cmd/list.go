package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hue_cmd"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		hue_cmd.List(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
