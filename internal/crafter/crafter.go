package crafter

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

type Crafter struct {
	// XIVCrafter settings
	amount int

	// Consumables
	food         string
	foodDuration time.Duration
	potion       string

	// In-game hotkeys
	confirm        string
	cancel         string
	macro1         string
	macro1Duration time.Duration
	macro2         string
	macro2Duration time.Duration
	macro3         string
	macro3Duration time.Duration

	// Helpers
	running         bool
	Paused          bool
	startTime       time.Time
	CurrentAmount   int
	foodCount       int
	foodStartTime   time.Time
	potionCount     int
	potionStartTime time.Time
}

// NewCrafter returns a pointer to a Crafter struct
func NewCrafter(startPause string, stop string, confirm string, cancel string) *Crafter {
	return &Crafter{
		// XIVCrafter settings
		amount: 1,

		// Consumables
		food:   "",
		potion: "",

		// In-game hotkeys
		confirm:        confirm,
		cancel:         cancel,
		macro1:         "",
		macro1Duration: 1 * time.Second,
		macro2:         "",
		macro2Duration: 1 * time.Second,
		macro3:         "",
		macro3Duration: 1 * time.Second,

		// Helpers
		running:         true,
		Paused:          true,
		startTime:       time.Time{},
		CurrentAmount:   0,
		foodCount:       0,
		foodStartTime:   time.Time{},
		potionCount:     0,
		potionStartTime: time.Time{},
	}
}

func (c *Crafter) SetRecipe(amount int, food string, foodDuration int, potion string, macro1 string, macro1Duration int, macro2 string, macro2Duration int, macro3 string, macro3Duration int) {
	// XIVCrafter settings
	c.amount = amount

	// Consumables
	c.food = food
	c.foodDuration = time.Duration(foodDuration) * time.Minute
	c.potion = potion

	// In-game hotkeys
	c.macro1 = macro1
	c.macro1Duration = time.Duration(macro1Duration) * time.Second
	c.macro2 = macro2
	c.macro2Duration = time.Duration(macro2Duration) * time.Second
	c.macro3 = macro3
	c.macro3Duration = time.Duration(macro3Duration) * time.Second
}

func (c *Crafter) ResetRecipe() {
	// Consumables
	c.food = ""
	c.foodCount = 0
	c.potion = ""
	c.potionCount = 0

	// In-game hotkeys
	c.macro1 = ""
	c.macro1Duration = 1
	c.macro2 = ""
	c.macro2Duration = 1
	c.macro3 = ""
	c.macro3Duration = 1

	// Helpers
	c.CurrentAmount = 0
	c.startTime = time.Time{}
	c.foodStartTime = time.Time{}
	c.potionStartTime = time.Time{}
}

func (c *Crafter) Run() {
	for {
		for c.running {
			// Main crafting loop
			for !c.Paused {
				// Get the start crafting time
				if c.startTime.IsZero() {
					c.startTime = time.Now()
				}

				c.startCraft()

				if c.food != "" {
					c.checkFood()
				}

				if c.potion != "" {
					c.checkPotion()
				}

				// Activate macro 1
				cobra.CheckErr(robotgo.KeyTap(c.macro1))
				time.Sleep(KeyDelay)
				time.Sleep(c.macro1Duration)

				if c.macro2 != "" {
					// Activate macro 2
					cobra.CheckErr(robotgo.KeyTap(c.macro2))
					time.Sleep(KeyDelay)
					time.Sleep(c.macro2Duration)
				}

				if c.macro3 != "" {
					// Activate macro 3
					cobra.CheckErr(robotgo.KeyTap(c.macro3))
					time.Sleep(KeyDelay)
					time.Sleep(c.macro3Duration)
				}

				c.CurrentAmount++
				if c.CurrentAmount >= c.amount {
					c.ExitProgram()
				}

				time.Sleep(EndCraftDelay)
			}

			time.Sleep(PauseDelay)
		}
	}
}

// StartProgram sets the paused value to false
func (c *Crafter) StartProgram() {
	c.Paused = false
}

// StopProgram sets the paused value to true
func (c *Crafter) StopProgram() {
	c.Paused = true
}

// ExitProgram sets the running value to false and the paused value to true
func (c *Crafter) ExitProgram() {
	c.running = false
	c.Paused = true
}

// startCraft sets up the crafting action
func (c *Crafter) startCraft() {
	cobra.CheckErr(robotgo.KeyTap(c.confirm))
	time.Sleep(KeyDelay)
	cobra.CheckErr(robotgo.KeyTap(c.confirm))
	time.Sleep(KeyDelay)
	cobra.CheckErr(robotgo.KeyTap(c.confirm))
	time.Sleep(KeyDelay)

	time.Sleep(StartCraftDelay)
}

// stopCraft closes the crafting action
func (c *Crafter) stopCraft() {
	cobra.CheckErr(robotgo.KeyTap(c.confirm))
	time.Sleep(KeyDelay)
	cobra.CheckErr(robotgo.KeyTap(c.cancel))
	time.Sleep(KeyDelay)
	cobra.CheckErr(robotgo.KeyTap(c.confirm))
	time.Sleep(KeyDelay)

	time.Sleep(EndCraftDelay)
}

// checkFood checks to see whether the food buff needs to be renewed
func (c *Crafter) checkFood() {
	if c.foodStartTime.IsZero() {
		c.consumeFood()
	} else {
		if time.Since(c.foodStartTime) > c.foodDuration {
			c.consumeFood()
		}
	}
}

// consumeFood renews the food buff
func (c *Crafter) consumeFood() {
	c.stopCraft()

	c.foodStartTime = time.Now()
	cobra.CheckErr(robotgo.KeyTap(c.food))
	c.foodCount++

	time.Sleep(EndCraftDelay)

	c.startCraft()
}

// checkPotion checks to see whether the potion buff needs to be renewed
func (c *Crafter) checkPotion() {
	if c.potionStartTime.IsZero() {
		c.consumePotion()
	} else {
		if time.Since(c.potionStartTime) > PotionDuration {
			c.consumePotion()
		}
	}
}

// consumePotion renews the potion buff
func (c *Crafter) consumePotion() {
	c.stopCraft()

	c.potionStartTime = time.Now()
	cobra.CheckErr(robotgo.KeyTap(c.potion))
	c.potionCount++

	time.Sleep(EndCraftDelay)

	c.startCraft()
}
