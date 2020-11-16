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
	Short: "fusion reactor control plane",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Powering off: Light %d ...\n", viper.GetInt(flagHueLightNumber2))
		return hue_cmd.Off(viper.GetString(flagHueIP), viper.GetString(flagHueUser), viper.GetInt(flagHueLightNumber2))
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
	offCmd.Flags().Int8P(flagHueLightNumber2, "l", 0, "Hue light number # 0 ... n")
	_ = viper.BindPFlag(flagHueLightNumber2, offCmd.Flags().Lookup(flagHueLightNumber2))
}
