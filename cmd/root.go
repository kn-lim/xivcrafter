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
		// VERBOSE, _ := cmd.PersistentFlags().GetBool("verbose")

		// Get Settings
		start_pause := viper.GetString("start_pause")
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
		var recipes []utils.Recipe
		if err := json.Unmarshal(recipesBytes, &recipes); err != nil {
			log.Fatalf("Unable to unmarshal recipes: %v", err)
		}

		// Create config
		config := utils.NewConfig()
		config.StartPause = start_pause
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

		// Setup UI
		items := []list.Item{
			tui.Item{ItemTitle: "Raspberry Pi’s", ItemDesc: "I have ’em all over my house"},
			tui.Item{ItemTitle: "Nutella", ItemDesc: "It's good on toast"},
			tui.Item{ItemTitle: "Bitter melon", ItemDesc: "It cools you down"},
			tui.Item{ItemTitle: "Nice socks", ItemDesc: "And by that I mean socks without holes"},
			tui.Item{ItemTitle: "Eight hours of sleep", ItemDesc: "I had this once"},
			tui.Item{ItemTitle: "Cats", ItemDesc: "Usually"},
			tui.Item{ItemTitle: "Plantasia, the album", ItemDesc: "My plants love it too"},
			tui.Item{ItemTitle: "Pour over coffee", ItemDesc: "It takes forever to make though"},
			tui.Item{ItemTitle: "VR", ItemDesc: "Virtual reality...what is there to say?"},
			tui.Item{ItemTitle: "Noguchi Lamps", ItemDesc: "Such pleasing organic forms"},
			tui.Item{ItemTitle: "Linux", ItemDesc: "Pretty much the best OS"},
			tui.Item{ItemTitle: "Business school", ItemDesc: "Just kidding"},
			tui.Item{ItemTitle: "Pottery", ItemDesc: "Wet clay is a great feeling"},
			tui.Item{ItemTitle: "Shampoo", ItemDesc: "Nothing like clean hair"},
			tui.Item{ItemTitle: "Table tennis", ItemDesc: "It’s surprisingly exhausting"},
			tui.Item{ItemTitle: "Milk crates", ItemDesc: "Great for packing in your extra stuff"},
			tui.Item{ItemTitle: "Afternoon tea", ItemDesc: "Especially the tea sandwich part"},
			tui.Item{ItemTitle: "Stickers", ItemDesc: "The thicker the vinyl the better"},
			tui.Item{ItemTitle: "20° Weather", ItemDesc: "Celsius, not Fahrenheit"},
			tui.Item{ItemTitle: "Warm light", ItemDesc: "Like around 2700 Kelvin"},
			tui.Item{ItemTitle: "The vernal equinox", ItemDesc: "The autumnal equinox is pretty good too"},
			tui.Item{ItemTitle: "Gaffer’s tape", ItemDesc: "Basically sticky fabric"},
			tui.Item{ItemTitle: "Terrycloth", ItemDesc: "In other words, towel fabric"},
		}

		m := tui.Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
		m.List.Title = "My Fave Things"

		p := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			log.Fatalf("Error running program: %v", err)
		}
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

	// Verbose
	// rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// XIVCrafter Hotkeys
	rootCmd.PersistentFlags().String("start-pause", "", "start/pause xivcrafter hotkey")
	rootCmd.PersistentFlags().String("stop", "", "stop xivcrafter hotkey")

	// In-Game Hotkeys
	rootCmd.PersistentFlags().String("confirm", "", "confirm hotkey")
	rootCmd.PersistentFlags().String("cancel", "", "cancel hotkey")

	// Viper Binds
	if err := viper.BindPFlag("start_pause", rootCmd.PersistentFlags().Lookup("start-pause")); err != nil {
		log.Fatalf("Error binding flag: %v", err)
	}

	if err := viper.BindPFlag("stop", rootCmd.PersistentFlags().Lookup("stop")); err != nil {
		log.Fatalf("Error binding flag: %v", err)
	}

	if err := viper.BindPFlag("confirm", rootCmd.PersistentFlags().Lookup("confirm")); err != nil {
		log.Fatalf("Error binding flag: %v", err)
	}

	if err := viper.BindPFlag("cancel", rootCmd.PersistentFlags().Lookup("cancel")); err != nil {
		log.Fatalf("Error binding flag: %v", err)
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

	viper.AutomaticEnv() // read in environment variables that match

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
