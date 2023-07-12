package cmd

import (
	"fmt"
	"os"

	hook "github.com/robotn/gohook"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kn-lim/xivcrafter/pkg/crafter"
	"github.com/kn-lim/xivcrafter/pkg/ui"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run"},
	Short:   "Starts XIVCrafter",
	Long:    `Starts XIVCrafter`,
	Run: func(cmd *cobra.Command, args []string) {
		VERBOSE, _ := rootCmd.PersistentFlags().GetBool("verbose")
		RANDOM, _ := rootCmd.PersistentFlags().GetBool("random")

		food := viper.GetString("food")
		foodDuration := viper.GetInt("foodDuration")
		potion := viper.GetString("potion")
		amount := viper.GetInt("amount")
		macro1 := viper.GetString("macro1")
		macro1Duration := viper.GetInt("macro1Duration")
		macro2 := viper.GetString("macro2")
		macro2Duration := viper.GetInt("macro2Duration")
		confirm := viper.GetString("confirm")
		cancel := viper.GetString("cancel")
		startPause := viper.GetString("startPause")
		stop := viper.GetString("stop")

		xiv := new(crafter.XIVCrafter)
		xiv.Init(food, foodDuration, potion, amount, macro1, macro1Duration, macro2, macro2Duration, confirm, cancel, startPause, stop)

		// Check if all keys and flags are valid
		if !(crafter.CheckKeys(*xiv) && crafter.CheckFlags(*xiv)) {
			os.Exit(1)
		}

		// Start UI
		ui := new(ui.UI)
		if !VERBOSE {
			ui.Init(xiv.MaxAmount)
		}

		go xiv.Run(ui, VERBOSE, RANDOM)

		hook.Register(hook.KeyDown, []string{xiv.StartPause}, func(e hook.Event) {
			if xiv.Running {
				if VERBOSE {
					fmt.Println("PAUSE XIVCRAFTER HOTKEY DETECTED")
				}

				xiv.StopProgram(ui, VERBOSE)
			} else {
				if VERBOSE {
					fmt.Println("START XIVCRAFTER HOTKEY DETECTED")
				}

				xiv.StartProgram(ui, VERBOSE)
			}
		})

		hook.Register(hook.KeyDown, []string{xiv.Stop}, func(e hook.Event) {
			if VERBOSE {
				fmt.Println("STOP XIVCRAFTER HOTKEY DETECTED")
			}

			xiv.ExitProgram(ui, VERBOSE)
			hook.StopEvent()
		})

		s := hook.Start()
		<-hook.Process(s)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
