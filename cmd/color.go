package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:           "color",
	Short:         "language agnostic eye pleasures",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("color called (unimplemented yet)")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)
}
