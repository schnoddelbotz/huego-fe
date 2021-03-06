package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         "prints a bestseller novel on-demand",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is huego-fe, version %s (%s)\n", Version, runtime.Version())
		fmt.Println("Visit https://github.com/schnoddelbotz/huego-fe for more information.")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
