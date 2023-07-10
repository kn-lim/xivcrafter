package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
		// Get Settings
		startPause := viper.GetString("start_pause")
		stop := viper.GetString("stop")
		confirm := viper.GetString("confirm")
		cancel := viper.GetString("cancel")

		// Read the 'recipes' field from the config
		recipesInterface := viper.Get("recipes")

		// Marshal the interface into JSON bytes
		recipesBytes, err := json.Marshal(recipesInterface)
		if err != nil {
			log.Fatalf("Unable to marshal recipes: %v", err)
		}

		// Unmarshal the JSON bytes into a slice of Recipe structs
		var recipes []tui.Recipe
		if err := json.Unmarshal(recipesBytes, &recipes); err != nil {
			log.Fatalf("Unable to unmarshal recipes: %v", err)
		}

		// Create config
		config := utils.NewConfig()
		config.StartPause = startPause
		config.Stop = stop
		config.Confirm = confirm
		config.Cancel = cancel
		config.Recipes = recipes

		// Validate Config
		// TODO

		// Setup Crafter
		// TODO

		// // Setup Start/Pause hotkey
		// hook.Register(hook.KeyDown, []string{start_pause}, func(e hook.Event) {
		// 	fmt.Println("Start/Pause")
		// })

		// // Setup Stop hotkey
		// hook.Register(hook.KeyDown, []string{stop}, func(e hook.Event) {
		// 	fmt.Println("Stop")
		// 	hook.End()
		// })

		// s := hook.Start()
		// <-hook.Process(s)

		// Setup Items for List
		items := []list.Item{}
		for _, recipe := range config.Recipes {
			items = append(items, tui.Recipe{
				Name:           recipe.Name,
				Food:           recipe.Food,
				FoodDuration:   recipe.FoodDuration,
				Potion:         recipe.Potion,
				Macro1:         recipe.Macro1,
				Macro1Duration: recipe.Macro1Duration,
				Macro2:         recipe.Macro2,
				Macro2Duration: recipe.Macro2Duration,
				Macro3:         recipe.Macro3,
				Macro3Duration: recipe.Macro3Duration,
			})
		}

		// Initialize Model
		m := tui.List{Recipes: list.New(items, tui.NewItemDelegate(), 0, 0)}
		m.Recipes.Title = "XIVCrafter"
		m.Recipes.Styles.Title = m.Recipes.Styles.Title.Padding(1, 3, 1).Bold(true).Background(tui.Primary).Foreground(tui.Tertiary)
		m.Recipes.SetShowHelp(false)

		// Run UI
		p := tea.NewProgram(m, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatalf("Error running program: %v", err)
		}

		// Return final crafting report
		// TODO
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
