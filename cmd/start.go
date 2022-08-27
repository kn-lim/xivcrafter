package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Constants
const DELAY = 2
const POTION_DURATION = 900

type XIVCrafter struct {
	running         bool
	programRunning  bool
	food            string
	foodDuration    int
	startFoodTime   int64
	potion          string
	startPotionTime int64
	currentAmount   int
	maxAmount       int
	macro1          string
	macro1Duration  int
	macro2          string
	macro2Duration  int
	confirm         string
	quit            string
	startPause      string
	stop            string
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run"},
	Short:   "Starts XIVCrafter",
	Long:    `Starts XIVCrafter`,
	Run: func(cmd *cobra.Command, args []string) {
		var VERBOSE bool
		VERBOSE, _ = rootCmd.PersistentFlags().GetBool("verbose")

		food := viper.GetString("food")
		foodDuration := viper.GetInt("foodDuration")
		potion := viper.GetString("potion")
		amount := viper.GetInt("amount")
		macro1 := viper.GetString("macro1")
		macro1Duration := viper.GetInt("macro1Duration")
		macro2 := viper.GetString("macro2")
		macro2Duration := viper.GetInt("macro2Duration")
		confirm := viper.GetString("confirm")
		quit := viper.GetString("quit")
		startPause := viper.GetString("startPause")
		stop := viper.GetString("stop")

		xiv := new(XIVCrafter)
		*xiv = create(food, foodDuration, potion, amount, macro1, macro1Duration, macro2, macro2Duration, confirm, quit, startPause, stop)

		go run(xiv, VERBOSE)

		hook.Register(hook.KeyDown, []string{xiv.startPause}, func(e hook.Event) {
			if xiv.running {
				if VERBOSE {
					fmt.Println("PAUSE XIVCRAFTER HOTKEY DETECTED")
				}

				stopProgram(xiv)
			} else {
				if VERBOSE {
					fmt.Println("START XIVCRAFTER HOTKEY DETECTED")
				}

				startProgram(xiv)
			}
		})

		hook.Register(hook.KeyDown, []string{xiv.stop}, func(e hook.Event) {
			if VERBOSE {
				fmt.Println("STOP XIVCRAFTER HOTKEY DETECTED")
			}

			exitProgram(xiv)
			hook.StopEvent()
		})

		s := hook.Start()
		<-hook.Process(s)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// Executes the crafter
func run(xiv *XIVCrafter, VERBOSE bool) {
	for xiv.programRunning {
		for xiv.running {
			startCraft(xiv, VERBOSE)

			if xiv.food != "" {
				checkFood(xiv, VERBOSE)
			}

			if xiv.potion != "" {
				checkPotion(xiv, VERBOSE)
			}

			// Activate Macro 1
			if VERBOSE {
				fmt.Println("ACTIVATING MACRO 1")
			}
			robotgo.KeyTap(xiv.macro1)
			delay(xiv.macro1Duration)

			// Activate Macro 2
			if xiv.macro2 != "" {
				if VERBOSE {
					fmt.Println("ACTIVATING MACRO 2")
				}

				robotgo.KeyTap(xiv.macro2)
				delay(xiv.macro2Duration)
			}

			xiv.currentAmount++

			s := fmt.Sprintf("CRAFTED: %d / %d", xiv.currentAmount, xiv.maxAmount)
			fmt.Println(s)

			if xiv.currentAmount >= xiv.maxAmount {
				exitProgram(xiv)
			}

			delay(DELAY)
		}

		delay(DELAY)
	}

	if VERBOSE {
		fmt.Println("XIVCRAFTER STOPPED")
	}

	// TODO: Find a cleaner way of doing this
	os.Exit(0)
}

// Helper Functions

func create(food string, foodDuration int, potion string, maxAmount int, macro1 string, macro1Duration int, macro2 string, macro2Duration int, confirm string, quit string, startPause string, stop string) XIVCrafter {
	// Convert Minutes to Seconds
	foodDurationSeconds := foodDuration * 60

	return XIVCrafter{
		running:         false,
		programRunning:  true,
		food:            strings.ToLower(food),
		foodDuration:    foodDurationSeconds,
		startFoodTime:   0,
		potion:          strings.ToLower(potion),
		startPotionTime: 0,
		currentAmount:   0,
		maxAmount:       maxAmount,
		macro1:          strings.ToLower(macro1),
		macro1Duration:  macro1Duration,
		macro2:          strings.ToLower(macro2),
		macro2Duration:  macro2Duration,
		confirm:         strings.ToLower(confirm),
		quit:            strings.ToLower(quit),
		startPause:      strings.ToLower(startPause),
		stop:            strings.ToLower(stop),
	}
}

func delay(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}

func startProgram(xiv *XIVCrafter) {
	fmt.Println("STARTING XIVCRAFTER")

	xiv.running = true
}

func stopProgram(xiv *XIVCrafter) {
	fmt.Println("STOPPING XIVCRAFTER")

	xiv.running = false
}

func exitProgram(xiv *XIVCrafter) {
	fmt.Println("EXITING XIVCRAFTER")

	xiv.running = false
	xiv.programRunning = false
}

func startCraft(xiv *XIVCrafter, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STARTING CRAFT")
	}

	robotgo.KeyTap(xiv.confirm)
	delay(DELAY)
	robotgo.KeyTap(xiv.confirm)
	delay(DELAY)
	robotgo.KeyTap(xiv.confirm)
	delay(DELAY)
}

func stopCraft(xiv *XIVCrafter, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STOPPING CRAFT")
	}

	robotgo.KeyTap(xiv.confirm)
	delay(DELAY)
	robotgo.KeyTap(xiv.quit)
	delay(DELAY)
	robotgo.KeyTap(xiv.confirm)
	delay(DELAY)
}

func checkFood(xiv *XIVCrafter, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("CHECKING FOOD")
	}

	if xiv.startFoodTime > 0 {
		difference := time.Now().Unix() - xiv.startFoodTime

		if difference > int64(xiv.foodDuration) {
			consumeFood(xiv, VERBOSE)
		}
	} else {
		consumeFood(xiv, VERBOSE)
	}
}

func consumeFood(xiv *XIVCrafter, VERBOSE bool) {
	stopCraft(xiv, VERBOSE)

	if VERBOSE {
		fmt.Println("CONSUMING FOOD")
	}

	xiv.startFoodTime = time.Now().Unix()
	robotgo.KeyTap(xiv.food)
	delay(DELAY)

	startCraft(xiv, VERBOSE)
}

func checkPotion(xiv *XIVCrafter, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("CHECKING POTION")
	}

	if xiv.startPotionTime > 0 {
		difference := time.Now().Unix() - xiv.startPotionTime

		if difference > POTION_DURATION {
			consumePotion(xiv, VERBOSE)
		}
	} else {
		consumePotion(xiv, VERBOSE)
	}
}

func consumePotion(xiv *XIVCrafter, VERBOSE bool) {
	stopCraft(xiv, VERBOSE)

	if VERBOSE {
		fmt.Println("CONSUMING POTION")
	}

	xiv.startPotionTime = time.Now().Unix()
	robotgo.KeyTap(xiv.potion)
	delay(DELAY)

	startCraft(xiv, VERBOSE)
}
