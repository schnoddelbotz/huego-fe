package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/schnoddelbotz/huego-fe/hue_cmd"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "BFG-9000 apparello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		hue_cmd.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
