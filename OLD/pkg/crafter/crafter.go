package crafter

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/kn-lim/xivcrafter/pkg/ui"
)

// Constants

const START_CRAFT_DELAY = 2
const END_CRAFT_DELAY = 3
const KEY_DELAY = 1
const POTION_DURATION = 900
const RANDOM_MAX = 5
const RANDOM_MIN = KEY_DELAY

// Global Variables

var RANDOM_DELAY bool = false
var START_TIME int64 = 0
var END_TIME int64 = 0

// XIVCrafter Struct
type XIVCrafter struct {
	Running         bool
	ProgramRunning  bool
	Food            string
	FoodCount       int
	FoodDuration    int
	StartFoodTime   int64
	Potion          string
	PotionCount     int
	StartPotionTime int64
	CurrentAmount   int
	MaxAmount       int
	Macro1          string
	Macro1Duration  int
	Macro2          string
	Macro2Duration  int
	Confirm         string
	Cancel          string
	StartPause      string
	Stop            string
}

// Init initializes the XIVCrafter struct with variables provided from flags
func (xiv *XIVCrafter) Init(food string, foodDuration int, potion string, maxAmount int, macro1 string, macro1Duration int, macro2 string, macro2Duration int, confirm string, cancel string, startPause string, stop string) {
	// Convert Minutes to Seconds
	foodDurationSeconds := foodDuration * 60

	*xiv = XIVCrafter{
		Running:         false,
		ProgramRunning:  true,
		Food:            strings.ToLower(food),
		FoodCount:       0,
		FoodDuration:    foodDurationSeconds,
		StartFoodTime:   0,
		Potion:          strings.ToLower(potion),
		PotionCount:     0,
		StartPotionTime: 0,
		CurrentAmount:   0,
		MaxAmount:       maxAmount,
		Macro1:          strings.ToLower(macro1),
		Macro1Duration:  macro1Duration,
		Macro2:          strings.ToLower(macro2),
		Macro2Duration:  macro2Duration,
		Confirm:         strings.ToLower(confirm),
		Cancel:          strings.ToLower(cancel),
		StartPause:      strings.ToLower(startPause),
		Stop:            strings.ToLower(stop),
	}
}

// Run provides the main logic to handle crafting
func (xiv *XIVCrafter) Run(ui *ui.UI, VERBOSE bool, RANDOM bool) {
	if RANDOM {
		if VERBOSE {
			fmt.Println("ADDING RANDOM DELAY")
		}

		RANDOM_DELAY = RANDOM

		rand.Seed(time.Now().UnixNano())
	}

	xiv.printControls()

	// Main crafting loop
	for xiv.ProgramRunning {
		// Set the crafting start time
		if START_TIME == 0 {
			START_TIME = time.Now().Unix()
		}

		// Start UI
		if !VERBOSE && xiv.CurrentAmount == 0 {
			ui.Start()
		}

		for xiv.Running {
			if !VERBOSE {
				ui.SetStart()
			}

			xiv.StartCraft(VERBOSE)

			if xiv.Food != "" {
				xiv.CheckFood(VERBOSE)
			}

			if xiv.Potion != "" {
				xiv.CheckPotion(VERBOSE)
			}

			// Activate Macro 1
			if VERBOSE {
				fmt.Println("ACTIVATING MACRO 1")
			}
			robotgo.KeyTap(xiv.Macro1)
			delay(KEY_DELAY)
			delay(xiv.Macro1Duration)

			// Activate Macro 2
			if xiv.Macro2 != "" {
				if VERBOSE {
					fmt.Println("ACTIVATING MACRO 2")
				}

				robotgo.KeyTap(xiv.Macro2)
				delay(KEY_DELAY)
				delay(xiv.Macro2Duration)
			}

			xiv.CurrentAmount++
			if xiv.CurrentAmount >= xiv.MaxAmount {
				xiv.ExitProgram(ui, VERBOSE)
			}

			if VERBOSE {
				s := fmt.Sprintf("CRAFTED: %d / %d", xiv.CurrentAmount, xiv.MaxAmount)
				fmt.Println(s)
			} else {
				ui.Increment()
			}

			if RANDOM_DELAY {
				delay(rand.Intn(RANDOM_MAX+END_CRAFT_DELAY) + END_CRAFT_DELAY)
			} else {
				delay(END_CRAFT_DELAY)
			}
		}

		if RANDOM_DELAY {
			delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
		} else {
			delay(KEY_DELAY)
		}

		if !xiv.Running && !VERBOSE && xiv.ProgramRunning {
			ui.SetStop()
		}
	}

	if VERBOSE {
		fmt.Println("XIVCRAFTER STOPPED")
	} else {
		fmt.Println()
	}

	// Set the craft ending time
	END_TIME = time.Now().Unix()

	// Print results
	xiv.result()

	os.Exit(0)
}

// StartProgram sets the Running value to true
func (xiv *XIVCrafter) StartProgram(ui *ui.UI, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STARTING XIVCRAFTER")
	} else {
		ui.SetStart()
	}

	xiv.Running = true
}

// StopProgram sets the Running value to false
func (xiv *XIVCrafter) StopProgram(ui *ui.UI, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STOPPING XIVCRAFTER")
	} else {
		ui.SetPause()
	}

	xiv.Running = false
}

// ExitProgram sets the Running and ProgramRunning value to false
func (xiv *XIVCrafter) ExitProgram(ui *ui.UI, VERBOSE bool) {
	if VERBOSE {
		fmt.Println("EXITING XIVCRAFTER")
	} else {
		ui.SetExit()
	}

	xiv.Running = false
	xiv.ProgramRunning = false
}

// StartCraft sets up the crafting action
func (xiv *XIVCrafter) StartCraft(VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STARTING CRAFT")
	}

	if RANDOM_DELAY {
		rand.Seed(time.Now().UnixNano())
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(KEY_DELAY)
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(KEY_DELAY)
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX+START_CRAFT_DELAY) + START_CRAFT_DELAY)
	} else {
		delay(START_CRAFT_DELAY)
	}
}

// StopCraft closes the crafting action
func (xiv *XIVCrafter) StopCraft(VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STOPPING CRAFT")
	}

	if RANDOM_DELAY {
		rand.Seed(time.Now().UnixNano())
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(KEY_DELAY)
	}

	robotgo.KeyTap(xiv.Cancel)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(KEY_DELAY)
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX+END_CRAFT_DELAY) + END_CRAFT_DELAY)
	} else {
		delay(END_CRAFT_DELAY)
	}
}

// CheckFood checks to see whether the food buff needs to be renewed
func (xiv *XIVCrafter) CheckFood(VERBOSE bool) {
	if VERBOSE {
		fmt.Println("CHECKING FOOD")
	}

	if xiv.StartFoodTime > 0 {
		difference := time.Now().Unix() - xiv.StartFoodTime

		if difference > int64(xiv.FoodDuration) {
			xiv.ConsumeFood(VERBOSE)
		}
	} else {
		xiv.ConsumeFood(VERBOSE)
	}
}

// ConsumeFood renews the food buff
func (xiv *XIVCrafter) ConsumeFood(VERBOSE bool) {
	xiv.StopCraft(VERBOSE)

	if VERBOSE {
		fmt.Println("CONSUMING FOOD")
	}

	if RANDOM_DELAY {
		rand.Seed(time.Now().UnixNano())
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(KEY_DELAY)
	}

	xiv.StartFoodTime = time.Now().Unix()
	robotgo.KeyTap(xiv.Food)
	xiv.FoodCount++

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX+END_CRAFT_DELAY) + END_CRAFT_DELAY)
	} else {
		delay(END_CRAFT_DELAY)
	}

	xiv.StartCraft(VERBOSE)
}

// CheckPotion checks to see whether the potion buff needs to be renewed
func (xiv *XIVCrafter) CheckPotion(VERBOSE bool) {
	if VERBOSE {
		fmt.Println("CHECKING POTION")
	}

	if xiv.StartPotionTime > 0 {
		difference := time.Now().Unix() - xiv.StartPotionTime

		if difference > POTION_DURATION {
			xiv.ConsumePotion(VERBOSE)
		}
	} else {
		xiv.ConsumePotion(VERBOSE)
	}
}

// ConsumePotion renews the potion buff
func (xiv *XIVCrafter) ConsumePotion(VERBOSE bool) {
	xiv.StopCraft(VERBOSE)

	if VERBOSE {
		fmt.Println("CONSUMING POTION")
	}

	if RANDOM_DELAY {
		rand.Seed(time.Now().UnixNano())
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(KEY_DELAY)
	}

	xiv.StartPotionTime = time.Now().Unix()
	robotgo.KeyTap(xiv.Potion)
	xiv.PotionCount++

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX+END_CRAFT_DELAY) + END_CRAFT_DELAY)
	} else {
		delay(END_CRAFT_DELAY)
	}

	xiv.StartCraft(VERBOSE)
}

// printControls prints out the program hotkeys
func (xiv *XIVCrafter) printControls() {
	s := fmt.Sprintf("Press \"%s\" to Start/Pause", xiv.StartPause)
	fmt.Println(s)

	s = fmt.Sprintf("Press \"%s\" to Stop", xiv.Stop)
	fmt.Println(s)

	fmt.Println()
}

// result prints out statistics
func (xiv *XIVCrafter) result() {
	fmt.Println("\nRESULTS:")

	s := fmt.Sprintf("CRAFTED: %d", xiv.CurrentAmount)
	fmt.Println(s)

	TIME_HOURS := ((END_TIME - START_TIME) / 60) / 60
	TIME_MINUTES := ((END_TIME - START_TIME) / 60) % 60
	TIME_SECONDS := (END_TIME - START_TIME) % 60
	if TIME_HOURS > 0 {
		s = fmt.Sprintf("TIME ELAPSED: %dhr %dmin %dsec", TIME_HOURS, TIME_MINUTES, TIME_SECONDS)
	} else if TIME_MINUTES == 0 {
		s = fmt.Sprintf("TIME ELAPSED: %dsec", TIME_SECONDS)
	} else {
		s = fmt.Sprintf("TIME ELAPSED: %dmin %dsec", TIME_MINUTES, TIME_SECONDS)
	}
	fmt.Println(s)

	if xiv.PotionCount > 0 {
		s = fmt.Sprintf("POTIONS USED: %d", xiv.PotionCount)
		fmt.Println(s)
	}

	if xiv.FoodCount > 0 {
		s = fmt.Sprintf("FOOD USED: %d", xiv.FoodCount)
		fmt.Println(s)
	}
}

// delay adds a delay
func delay(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}
