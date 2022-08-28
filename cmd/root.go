package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xivcrafter",
	Short: "A FFXIV Automated Crafting Tool",
	Long:  `Automatically activates multiple crafting macros while refreshing food and potion buffs.`,
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME\\.xivcrafter.yaml)")

	// Verbose
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// Flags
	rootCmd.PersistentFlags().String("food", "", "food hotkey")
	rootCmd.PersistentFlags().Int("foodDuration", 30, "food duration (30/40/45)")
	rootCmd.PersistentFlags().String("potion", "", "potion hotkey")
	rootCmd.PersistentFlags().Int("amount", 0, "amount to craft")
	rootCmd.PersistentFlags().String("macro1", "", "macro 1 hotkey")
	rootCmd.PersistentFlags().Int("macro1Duration", 0, "macro 1 duration")
	rootCmd.PersistentFlags().String("macro2", "", "macro 2 hotkey")
	rootCmd.PersistentFlags().Int("macro2Duration", 0, "macro 2 duration")
	rootCmd.PersistentFlags().String("confirm", "", "confirm hotkey")
	rootCmd.PersistentFlags().String("cancel", "", "cancel hotkey")
	rootCmd.PersistentFlags().String("startPause", "", "start/pause xivcrafter hotkey")
	rootCmd.PersistentFlags().String("stop", "", "stop xivcrafter hotkey")

	// Viper
	viper.BindPFlag("food", rootCmd.PersistentFlags().Lookup("food"))
	viper.BindPFlag("foodDuration", rootCmd.PersistentFlags().Lookup("foodDuration"))
	viper.BindPFlag("potion", rootCmd.PersistentFlags().Lookup("potion"))
	viper.BindPFlag("amount", rootCmd.PersistentFlags().Lookup("amount"))
	viper.BindPFlag("macro1", rootCmd.PersistentFlags().Lookup("macro1"))
	viper.BindPFlag("macro1Duration", rootCmd.PersistentFlags().Lookup("macro1Duration"))
	viper.BindPFlag("macro2", rootCmd.PersistentFlags().Lookup("macro2"))
	viper.BindPFlag("macro2Duration", rootCmd.PersistentFlags().Lookup("macro2Duration"))
	viper.BindPFlag("confirm", rootCmd.PersistentFlags().Lookup("confirm"))
	viper.BindPFlag("cancel", rootCmd.PersistentFlags().Lookup("cancel"))
	viper.BindPFlag("startPause", rootCmd.PersistentFlags().Lookup("startPause"))
	viper.BindPFlag("stop", rootCmd.PersistentFlags().Lookup("stop"))
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
		viper.SetConfigType("yaml")
		viper.SetConfigName(".xivcrafter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
