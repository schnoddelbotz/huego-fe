package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// brightnessCmd represents the brightness command
var brightnessCmd = &cobra.Command{
	Use:   "brightness",
	Short: "control gravity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brightness called")
	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)
}
