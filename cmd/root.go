/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kn-lim/xivcrafter/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xivcrafter",
	Short: "A FFXIV Automated Crafting Tool",
	Long:  `Automatically activates multiple crafting macros while refreshing food and potion buffs.`,

	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.xivcrafter.json)")

	// Verbose
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// XIVCrafter Hotkeys
	rootCmd.PersistentFlags().String("startPause", "", "start/pause xivcrafter hotkey")
	rootCmd.PersistentFlags().String("stop", "", "stop xivcrafter hotkey")

	// In-Game Hotkeys
	rootCmd.PersistentFlags().String("confirm", "", "confirm hotkey")
	rootCmd.PersistentFlags().String("cancel", "", "cancel hotkey")

	// Viper Binds
	viper.BindPFlag("start_pause", rootCmd.PersistentFlags().Lookup("startPause"))
	viper.BindPFlag("stop", rootCmd.PersistentFlags().Lookup("stop"))
	viper.BindPFlag("confirm", rootCmd.PersistentFlags().Lookup("confirm"))
	viper.BindPFlag("cancel", rootCmd.PersistentFlags().Lookup("cancel"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".xivcrafter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".xivcrafter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		// Create config file in home directory
		config := utils.NewConfig()

		// Convert to JSON
		data, err := json.MarshalIndent(config, "", "  ")
		cobra.CheckErr(err)

		// Write to JSON File
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		file := home + "/.xivcrafter.json"
		err = ioutil.WriteFile(file, data, 0644)
		cobra.CheckErr(err)

		// Set new config file
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Creating new config file:", viper.ConfigFileUsed())
		}
	}
}
