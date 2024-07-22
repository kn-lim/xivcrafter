package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kn-lim/xivcrafter/internal/crafter"
	"github.com/kn-lim/xivcrafter/internal/tui"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xivcrafter",
	Short: "A FFXIV Automated Crafting Tool",
	Long:  `Automatically activates multiple crafting macros while refreshing food and potion buffs.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Set debug mode
		debug, err := cmd.PersistentFlags().GetBool("debug")
		if err != nil {
			cobra.CheckErr(err)
		}

		if debug {
			// Setup debug log file
			homeDir, err := os.UserHomeDir()
			if err != nil {
				cobra.CheckErr(err)
			}
			path := homeDir + "/.xivcrafter-log.jsonl"
			f, err := tea.LogToFile(path, "debug")
			if err != nil {
				cobra.CheckErr(err)
			}
			defer f.Close()

			// Setup Logger
			utils.Logger, err = utils.CreateLogger(path)
			if err != nil {
				cobra.CheckErr(err)
			}
		}

		// Get config path
		utils.ConfigPath = viper.ConfigFileUsed()

		// Get settings
		delay := viper.GetInt("delay")
		keyDelay := viper.GetInt("key_delay")
		startPause := viper.GetString("start_pause")
		stop := viper.GetString("stop")
		confirm := viper.GetString("confirm")
		cancel := viper.GetString("cancel")

		utils.Log("Infow", "using xivcrafter settings",
			"delay", delay,
			"key_delay", keyDelay,
			"start_pause", startPause,
			"stop", stop,
			"confirm", confirm,
			"cancel", cancel,
		)

		// Read the 'recipes' field from the config
		recipesInterface := viper.Get("recipes")

		// Marshal the interface into JSON bytes
		recipesBytes, err := json.Marshal(recipesInterface)
		cobra.CheckErr(err)

		// Unmarshal the JSON bytes into a slice of Recipe structs
		var recipes []utils.Recipe
		cobra.CheckErr(json.Unmarshal(recipesBytes, &recipes))

		utils.Log("Infow", "loaded recipes",
			"count", len(recipes),
		)

		// Setup UI
		tui.StartPause = startPause
		tui.Stop = stop
		tui.Confirm = confirm
		tui.Cancel = cancel

		// Setup Items for List model
		items := []list.Item{}
		if len(recipes) != 1 || recipes[0].Name != "" { // check for new config
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

		// Setup UpdateSettings model
		tui.Models[tui.ChangeSettings] = tui.NewUpdateSettings()

		// Setup List model
		tui.Models[tui.Recipes] = tui.NewList(items)

		// Setup UpdateRecipe model
		tui.Models[tui.ChangeRecipe] = tui.NewUpdateRecipe()

		// Setup Input model
		tui.Models[tui.Amount] = tui.NewInput()

		// Setup Progress model
		tui.Models[tui.Crafter] = tui.NewProgress()

		var p *tea.Program

		// Check if XIVCrafter settings are valid
		if err := utils.ValidateSettings(startPause, stop, confirm, cancel); err != nil {
			utils.Log("Errorw", "xivcrafter settings invalid",
				"error", err,
			)

			// Show error message
			model := tui.Models[tui.ChangeSettings].(*tui.UpdateSettings)
			model.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render(err.Error())
			model.AddPlaceholders(*tui.NewSettings(startPause, stop, confirm, cancel))
			tui.Models[tui.ChangeSettings] = model

			p = tea.NewProgram(tui.Models[tui.ChangeSettings], tea.WithAltScreen())
		} else {
			// Check if the entire config is valid
			errResult, err := utils.Validate(startPause, stop, confirm, cancel, tui.ConvertItemsToRecipes(tui.ConvertListItemToItem(items)))
			if err != nil {
				utils.Log("Errorw", "invalid recipe",
					"recipe", errResult,
				)

				// Get recipe with error
				var errorItem tui.Item
				for _, item := range items {
					if errResult == item.(tui.Item).Name {
						errorItem = item.(tui.Item)
						break
					}
				}

				// Show error message
				updateRecipeModel := tui.Models[tui.ChangeRecipe].(*tui.UpdateRecipe)
				updateRecipeModel.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render(err.Error())
				updateRecipeModel.AddPlaceholders(errorItem)
				tui.Models[tui.ChangeRecipe] = updateRecipeModel

				// Tell List model to replace recipe
				listModel := tui.Models[tui.Recipes].(*tui.List)
				listModel.Edit = true
				tui.Models[tui.Recipes] = listModel

				p = tea.NewProgram(tui.Models[tui.ChangeRecipe], tea.WithAltScreen())
			} else {
				p = tea.NewProgram(tui.Models[tui.Recipes], tea.WithAltScreen())
			}
		}

		// Run UI
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
	rootCmd.PersistentFlags().Bool("debug", false, "enable debugging (debug log location is $HOME/.xivcrafter-log.jsonl)")

	// XIVCrafter Settings
	rootCmd.PersistentFlags().Int("delay", 0, "crafting delay in milliseconds")
	rootCmd.PersistentFlags().Int("key-delay", 0, "keyboard input delay in milliseconds")

	// XIVCrafter Hotkeys
	rootCmd.PersistentFlags().String("start-pause", "", "start/pause xivcrafter hotkey")
	rootCmd.PersistentFlags().String("stop", "", "stop xivcrafter hotkey")

	// In-Game Hotkeys
	rootCmd.PersistentFlags().String("confirm", "", "confirm hotkey")
	rootCmd.PersistentFlags().String("cancel", "", "cancel hotkey")

	// Viper Binds
	if err := viper.BindPFlag("delay", rootCmd.PersistentFlags().Lookup("delay")); err != nil {
		cobra.CheckErr(err)
	}
	if err := viper.BindPFlag("key_delay", rootCmd.PersistentFlags().Lookup("key-delay")); err != nil {
		cobra.CheckErr(err)
	}
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
		err = os.WriteFile(file, data, 0o644)
		cobra.CheckErr(err)

		// Set new config file
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			cobra.CheckErr(err)
		}
	}
}
