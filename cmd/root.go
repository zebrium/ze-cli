// Package cmd Copyright © 2024 ScienceLogic Inc
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var version = "2.2.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ze",
	Short: "Zebrium CLI",
	Long: `ze is a CLI library for Zebrium. This application allows uses to submit local files to Zebrium.  
This also provides support for interacting with Batch bundles`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ze)")
	rootCmd.PersistentFlags().StringP("url", "u", "https://cloud.zebrium.com", "zapi endpoint.")
	rootCmd.PersistentFlags().StringP("auth", "a", "", "Zebrium authentication token.  Can be found under Integrations & Collectors in the Zebrium UI ")
	rootCmd.PersistentFlags().StringP("api", "t", "", "Zebrium API token.  Can be found under Access Tokens in the Zebrium UI")
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if len(cfgFile) != 0 {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")

	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ze" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ze")
	}
	viper.SetEnvPrefix("ze")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
