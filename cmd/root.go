package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

const (
	flagHueUser         = "hue-user"
	flagHueIP           = "hue-ip"
	flagHueLightNumber1 = "hue-light-on"
	flagHueLightNumber2 = "hue-light-off"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "huego-fe",
	Short: "huego-fe can control your philips hue stuff",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.huego-fe.yaml)")
	rootCmd.PersistentFlags().StringP(flagHueUser, "u", "", "Hue bridge user, see: huego-fe login -h")
	rootCmd.PersistentFlags().StringP(flagHueIP, "i", "", "Hue bridge IP, see: huego-fe login -h")
	// Cobra also supports local flags, which will only run when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = viper.BindPFlag(flagHueUser, rootCmd.PersistentFlags().Lookup(flagHueUser))
	_ = viper.BindPFlag(flagHueIP, rootCmd.PersistentFlags().Lookup(flagHueIP))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".huego-fe" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".huego-fe")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
