package crafter

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

// Constants
const DELAY = 2
const POTION_DURATION = 900

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

func (xiv *XIVCrafter) Run(VERBOSE bool) {
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

func (xiv *XIVCrafter) StartProgram() {
	fmt.Println("STARTING XIVCRAFTER")

	xiv.Running = true
}

func (xiv *XIVCrafter) StopProgram() {
	fmt.Println("STOPPING XIVCRAFTER")

	xiv.Running = false
}

func (xiv *XIVCrafter) ExitProgram() {
	fmt.Println("EXITING XIVCRAFTER")

	xiv.Running = false
	xiv.ProgramRunning = false
}

func (xiv *XIVCrafter) StartCraft(VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STARTING CRAFT")
	}

	robotgo.KeyTap(xiv.Confirm)
	delay(DELAY)
	robotgo.KeyTap(xiv.Confirm)
	delay(DELAY)
	robotgo.KeyTap(xiv.Confirm)
	delay(DELAY)
}

func (xiv *XIVCrafter) StopCraft(VERBOSE bool) {
	if VERBOSE {
		fmt.Println("STOPPING CRAFT")
	}

	robotgo.KeyTap(xiv.Confirm)
	delay(DELAY)
	robotgo.KeyTap(xiv.Cancel)
	delay(DELAY)
	robotgo.KeyTap(xiv.Confirm)
	delay(DELAY)
}

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

func (xiv *XIVCrafter) ConsumeFood(VERBOSE bool) {
	xiv.StopCraft(VERBOSE)

	if VERBOSE {
		fmt.Println("CONSUMING FOOD")
	}

	xiv.StartFoodTime = time.Now().Unix()
	robotgo.KeyTap(xiv.Food)
	delay(DELAY)

	xiv.StartCraft(VERBOSE)
}

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

func (xiv *XIVCrafter) ConsumePotion(VERBOSE bool) {
	xiv.StopCraft(VERBOSE)

	if VERBOSE {
		fmt.Println("CONSUMING POTION")
	}

	xiv.StartPotionTime = time.Now().Unix()
	robotgo.KeyTap(xiv.Potion)
	delay(DELAY)

	xiv.StartCraft(VERBOSE)
}

// Helper Functions
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

func delay(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}