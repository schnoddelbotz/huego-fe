package cmd

import (
	"github.com/schnoddelbotz/huego-fe/hue_cmd"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l"},
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// todo: how-to set both args required down here...? :/
		hue_cmd.List(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
