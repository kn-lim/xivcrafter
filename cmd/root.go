package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kn-lim/xivcrafter/internal/crafter"
	"github.com/kn-lim/xivcrafter/internal/tui"
	"github.com/kn-lim/xivcrafter/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xivcrafter",
	Short: "A FFXIV Automated Crafting Tool",
	Long:  `Automatically activates multiple crafting macros while refreshing food and potion buffs.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Debug mode
		Debug, _ := cmd.PersistentFlags().GetBool("debug")
		if Debug {
			home, _ := os.UserHomeDir()
			path := home + "/.xivcrafter-debug.log"
			f, err := tea.LogToFile(path, "debug")
			cobra.CheckErr(err)
			defer f.Close()

			utils.Logger = log.New(f, "", log.LstdFlags)
		}

		// Get config path
		utils.ConfigPath = viper.ConfigFileUsed()

		// Get settings
		startPause := viper.GetString("start_pause")
		stop := viper.GetString("stop")
		confirm := viper.GetString("confirm")
		cancel := viper.GetString("cancel")

		if utils.Logger != nil {
			utils.Logger.Printf("StartPause: %s, Stop: %s, Confirm: %s, Cancel: %s\n", startPause, stop, confirm, cancel)
		}

		// Read the 'recipes' field from the config
		recipesInterface := viper.Get("recipes")

		// Marshal the interface into JSON bytes
		recipesBytes, err := json.Marshal(recipesInterface)
		cobra.CheckErr(err)

		// Unmarshal the JSON bytes into a slice of Recipe structs
		var recipes []utils.Recipe
		cobra.CheckErr(json.Unmarshal(recipesBytes, &recipes))

		if utils.Logger != nil {
			utils.Logger.Printf("Number of Recipes: %v\n", len(recipes))
		}

		// Validate Config
		// TODO

		// Setup Items for List model
		items := []list.Item{}
		if len(recipes) != 1 || recipes[0].Name != "" {
			for _, recipe := range recipes {
				items = append(items, tui.Item{
					Name:           recipe.Name,
					Food:           strings.ToLower(recipe.Food),
					FoodDuration:   recipe.FoodDuration,
					Potion:         strings.ToLower(recipe.Potion),
					Macro1:         strings.ToLower(recipe.Macro1),
					Macro1Duration: recipe.Macro1Duration,
					Macro2:         strings.ToLower(recipe.Macro2),
					Macro2Duration: recipe.Macro2Duration,
					Macro3:         strings.ToLower(recipe.Macro3),
					Macro3Duration: recipe.Macro3Duration,
				})
			}
		} // else will return no items, as this will indicate a new config

		// Setup List model
		tui.Models[tui.Recipes] = tui.NewList(startPause, stop, confirm, cancel, items)

		// Setup Update model
		tui.Models[tui.UpdateRecipe] = tui.NewUpdate()

		// Setup Input model
		tui.Models[tui.Amount] = tui.NewInput()

		// Setup Progress model
		tui.Models[tui.Crafter] = tui.NewProgress(startPause, stop, confirm, cancel)

		// Run UI
		p := tea.NewProgram(tui.Models[tui.Recipes], tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			cobra.CheckErr(err)
		}

		// Return final crafting report
		crafter.PrintResults()
	},
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

	// Debug
	rootCmd.PersistentFlags().Bool("debug", false, "enable debugging (debug log location is $HOME/.xivcrafter-debug.log)")

	// XIVCrafter Hotkeys
	rootCmd.PersistentFlags().String("start-pause", "", "start/pause xivcrafter hotkey")
	rootCmd.PersistentFlags().String("stop", "", "stop xivcrafter hotkey")

	// In-Game Hotkeys
	rootCmd.PersistentFlags().String("confirm", "", "confirm hotkey")
	rootCmd.PersistentFlags().String("cancel", "", "cancel hotkey")

	// Viper Binds
	if err := viper.BindPFlag("start_pause", rootCmd.PersistentFlags().Lookup("start-pause")); err != nil {
		cobra.CheckErr(err)
	}

	if err := viper.BindPFlag("stop", rootCmd.PersistentFlags().Lookup("stop")); err != nil {
		cobra.CheckErr(err)
	}

	if err := viper.BindPFlag("confirm", rootCmd.PersistentFlags().Lookup("confirm")); err != nil {
		cobra.CheckErr(err)
	}

	if err := viper.BindPFlag("cancel", rootCmd.PersistentFlags().Lookup("cancel")); err != nil {
		cobra.CheckErr(err)
	}
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

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Create config file in home directory
		config := utils.NewConfig()

		// Convert to JSON
		data, err := json.MarshalIndent(config, "", "  ")
		cobra.CheckErr(err)

		// Write to JSON File
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		file := home + "/.xivcrafter.json"
		err = os.WriteFile(file, data, 0644)
		cobra.CheckErr(err)

		// Set new config file
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error creating file: %s", file)
		}
	}
}
