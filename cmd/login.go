package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/hueController"
)

var loginCmd = &cobra.Command{
	Use:           "login",
	Short:         "Discover Hue bridge and log in -- press link button first!",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ip, user, err := hueController.Login()
		if err != nil {
			return err
		}
		viper.Set(flagHueIP, ip)
		viper.Set(flagHueUser, user)
		err = viper.WriteConfig()
		if err != nil {
			return err
		}
		fmt.Printf("Login succes, saved to: %s\n", viper.ConfigFileUsed())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
