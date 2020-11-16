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
	Short: "engage rocket launcher",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Powering on: Light %d ...\n", viper.GetInt(flagHueLightNumber1))
		return hue_cmd.On(viper.GetString(flagHueIP), viper.GetString(flagHueUser), viper.GetInt(flagHueLightNumber1))
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.Flags().Int8P(flagHueLightNumber1, "l", 0, "Hue light number # 0 ... n")
	_ = viper.BindPFlag(flagHueLightNumber1, onCmd.Flags().Lookup(flagHueLightNumber1))
}
