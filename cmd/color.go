package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "language agnostic eye pleasures",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("color called")
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)
}
