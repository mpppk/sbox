package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var projectName string
var serverName string

var RootCmd = &cobra.Command{
	Use:   "sbox",
	Short: "CLI for Scrapbox",
	Long:  ``,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/sbox/.sbox.yaml)")
	RootCmd.PersistentFlags().StringVar(&projectName, "project", "", "target project")
	viper.BindPFlag("project", RootCmd.PersistentFlags().Lookup("project"))
	RootCmd.PersistentFlags().StringVar(&serverName, "server", "http://scrapbox.io", "target server")
	viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".sbox") // name of config file (without extension)
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "sbox")
	viper.AddConfigPath(configPath) // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
