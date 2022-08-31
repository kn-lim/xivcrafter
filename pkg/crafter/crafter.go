package crafter

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

// Constants
const DELAY = 2
const POTION_DURATION = 900
const RANDOM_MAX = 3
const RANDOM_MIN = 1

// Global Variables
var RANDOM_DELAY bool

// XIVCrafter Struct
type XIVCrafter struct {
	Running         bool
	ProgramRunning  bool
	Food            string
	FoodDuration    int
	StartFoodTime   int64
	Potion          string
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

// Run provides the main logic to handle crafting
func (xiv *XIVCrafter) Run(VERBOSE bool, RANDOM bool) {
	if RANDOM {
		if VERBOSE {
			fmt.Println("ADDING RANDOM DELAY")
		}

		RANDOM_DELAY = RANDOM

		rand.Seed(time.Now().UnixNano())
	}

	for xiv.ProgramRunning {
		for xiv.Running {
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
			delay(xiv.Macro1Duration)

			// Activate Macro 2
			if xiv.Macro2 != "" {
				if VERBOSE {
					fmt.Println("ACTIVATING MACRO 2")
				}

				robotgo.KeyTap(xiv.Macro2)
				delay(xiv.Macro2Duration)
			}

			xiv.CurrentAmount++

			s := fmt.Sprintf("CRAFTED: %d / %d", xiv.CurrentAmount, xiv.MaxAmount)
			fmt.Println(s)

			if xiv.CurrentAmount >= xiv.MaxAmount {
				xiv.ExitProgram()
			}

			if RANDOM_DELAY {
				delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
			} else {
				delay(DELAY)
			}
		}

		if RANDOM_DELAY {
			delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
		} else {
			delay(DELAY)
		}
	}

	if VERBOSE {
		fmt.Println("XIVCRAFTER STOPPED")
	}

	// TODO: Find a cleaner way of doing this
	os.Exit(0)
}

// StartProgram sets the Running value to true
func (xiv *XIVCrafter) StartProgram() {
	fmt.Println("STARTING XIVCRAFTER")

	xiv.Running = true
}

// StopProgram sets the Running value to false
func (xiv *XIVCrafter) StopProgram() {
	fmt.Println("STOPPING XIVCRAFTER")

	xiv.Running = false
}

// ExitProgram sets the Running and ProgramRunning value to false
func (xiv *XIVCrafter) ExitProgram() {
	fmt.Println("EXITING XIVCRAFTER")

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
		delay(DELAY)
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(DELAY)
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(DELAY)
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
		delay(DELAY)
	}

	robotgo.KeyTap(xiv.Cancel)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(DELAY)
	}

	robotgo.KeyTap(xiv.Confirm)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(DELAY)
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
	}

	xiv.StartFoodTime = time.Now().Unix()
	robotgo.KeyTap(xiv.Food)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(DELAY)
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
	}

	xiv.StartPotionTime = time.Now().Unix()
	robotgo.KeyTap(xiv.Potion)

	if RANDOM_DELAY {
		delay(rand.Intn(RANDOM_MAX) + RANDOM_MIN)
	} else {
		delay(DELAY)
	}

	xiv.StartCraft(VERBOSE)
}

// Init initializes the XIVCrafter struct with variables provided from flags
func (xiv *XIVCrafter) Init(food string, foodDuration int, potion string, maxAmount int, macro1 string, macro1Duration int, macro2 string, macro2Duration int, confirm string, cancel string, startPause string, stop string) {
	// Convert Minutes to Seconds
	foodDurationSeconds := foodDuration * 60

	*xiv = XIVCrafter{
		Running:         false,
		ProgramRunning:  true,
		Food:            strings.ToLower(food),
		FoodDuration:    foodDurationSeconds,
		StartFoodTime:   0,
		Potion:          strings.ToLower(potion),
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

// delay adds a delay
func delay(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}
