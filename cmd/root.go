package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/schnoddelbotz/huego-fe/gui"
	"github.com/schnoddelbotz/huego-fe/hueController"
)

var Version = "0.0.0-dev"
var cfgFile string

const (
	flagHueUser  = "hue-user"
	flagHueIP    = "hue-ip"
	flagHueLight = "hue-light"
)

var rootCmd = &cobra.Command{
	Use:   "huego-fe",
	Short: "huego-fe can control your philips hue stuff",
	Run: func(cmd *cobra.Command, args []string) {
		// start UI if huego-fe called w/o args
		controller := hueController.New(viper.GetString(flagHueIP), viper.GetString(flagHueUser))
		lightId := viper.GetInt(flagHueLight)
		gui.Main(controller, lightId)
	},
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
	// should disable moustrap on win (would display help text in console if clicked via explorer)
	cobra.MousetrapHelpText = ""
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.huego-fe.yaml)")
	rootCmd.PersistentFlags().StringP(flagHueUser, "u", "", "Hue bridge user/token [$HUE_USER], see: huego-fe login -h")
	rootCmd.PersistentFlags().StringP(flagHueIP, "i", "", "Hue bridge IP [$HUE_IP] , see: huego-fe login -h")
	rootCmd.PersistentFlags().IntP(flagHueLight, "l", 1, "Hue light No.# [$HUE_LIGHT], see: huego-fe list")
	// make flags like --hue-ip available to app if HUE_IP in env:
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("yml")
	_ = viper.BindPFlag(flagHueUser, rootCmd.PersistentFlags().Lookup(flagHueUser))
	_ = viper.BindPFlag(flagHueIP, rootCmd.PersistentFlags().Lookup(flagHueIP))
	_ = viper.BindPFlag(flagHueLight, rootCmd.PersistentFlags().Lookup(flagHueLight))
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
		// todo: maybe -v / --debug?
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
