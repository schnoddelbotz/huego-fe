package cmd

import (
	"github.com/spf13/cobra"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "BFG-9000 apparello",
	Run: func(cmd *cobra.Command, args []string) {
		hueController.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
