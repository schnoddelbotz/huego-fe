package cmd

import (
	"fmt"
	"github.com/schnoddelbotz/huego-fe/hue_cmd"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Aliases: []string{"0"},
	Short: "fusion reactor control plane",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Powering off: Light %d ...\n", viper.GetInt(flagHueLight))
		return hue_cmd.Off(viper.GetString(flagHueIP), viper.GetString(flagHueUser), viper.GetInt(flagHueLight))
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
