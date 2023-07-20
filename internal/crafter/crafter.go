package crafter

import (
	"context"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-vgo/robotgo"
	"github.com/kn-lim/xivcrafter/internal/utils"
	hook "github.com/robotn/gohook"
	"github.com/spf13/cobra"
)

type Crafter struct {
	// XIVCrafter settings
	amount     int
	startPause string
	stop       string

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

	// Information
	startTime       time.Time
	CurrentAmount   int
	FoodCount       int
	foodStartTime   time.Time
	PotionCount     int
	potionStartTime time.Time

	// Helpers
	running               bool
	paused                bool
	exitOnce              sync.Once
	StopCrafterContext    context.Context
	StopCrafterCancelFunc context.CancelFunc
	StopHookContext       context.Context
	StopHookCancelFunc    context.CancelFunc
}

// NewCrafter returns a pointer to a Crafter struct
func NewCrafter(startPause string, stop string, confirm string, cancel string) *Crafter {
	crafterCtx, crafterCancelFunc := context.WithCancel(context.Background())
	hookCtx, hookCancelFunc := context.WithCancel(context.Background())

	return &Crafter{
		// XIVCrafter settings
		amount:     1,
		startPause: startPause,
		stop:       stop,

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

		// Information
		startTime:       time.Time{},
		CurrentAmount:   0,
		FoodCount:       0,
		foodStartTime:   time.Time{},
		PotionCount:     0,
		potionStartTime: time.Time{},

		// Helpers
		running:               true,
		paused:                true,
		StopCrafterContext:    crafterCtx,
		StopCrafterCancelFunc: crafterCancelFunc,
		StopHookContext:       hookCtx,
		StopHookCancelFunc:    hookCancelFunc,
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
	c.potion = ""

	// In-game hotkeys
	c.macro1 = ""
	c.macro1Duration = 1
	c.macro2 = ""
	c.macro2Duration = 1
	c.macro3 = ""
	c.macro3Duration = 1

	// Information
	c.startTime = time.Time{}
	c.CurrentAmount = 0
	c.FoodCount = 0
	c.foodStartTime = time.Time{}
	c.PotionCount = 0
	c.potionStartTime = time.Time{}
}

func (c *Crafter) Run() tea.Cmd {
	return func() tea.Msg {
		go func() {
			for c.running {
				// Main crafting loop
				for !c.paused {
					select {
					case <-c.StopCrafterContext.Done():
						if utils.Logger != nil {
							utils.Logger.Println("Stopping crafter context")
						}

						c.StopCrafterCancelFunc()
						return
					default:
						if utils.Logger != nil {
							utils.Logger.Printf("Starting craft %v / %v\n", c.CurrentAmount, c.amount)
						}

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

						if utils.Logger != nil {
							utils.Logger.Printf("Finishing craft %v / %v\n", c.CurrentAmount, c.amount)
						}

						time.Sleep(EndCraftDelay)
					}
				}

				time.Sleep(PauseDelay)
			}
		}()

		return nil
	}
}

// StartProgram sets the paused value to false
func (c *Crafter) StartProgram() {
	if utils.Logger != nil {
		utils.Logger.Println("Starting XIVCrafter")
	}

	c.paused = false
}

// StopProgram sets the paused value to true
func (c *Crafter) StopProgram() {
	if utils.Logger != nil {
		utils.Logger.Println("Pausing XIVCrafter")
	}

	c.paused = true
}

// ExitProgram sets the running value to false and the paused value to true
func (c *Crafter) ExitProgram() {
	c.exitOnce.Do(func() {
		if utils.Logger != nil {
			utils.Logger.Println("Exiting XIVCrafter")
		}

		c.StopCrafterCancelFunc()
		c.StopHookCancelFunc()

		c.running = false
		c.paused = true
	})
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
	if utils.Logger != nil {
		utils.Logger.Println("Consuming food")
	}

	c.stopCraft()

	c.foodStartTime = time.Now()
	cobra.CheckErr(robotgo.KeyTap(c.food))
	c.FoodCount++

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
	if utils.Logger != nil {
		utils.Logger.Println("Consuming potion")
	}

	c.stopCraft()

	c.potionStartTime = time.Now()
	cobra.CheckErr(robotgo.KeyTap(c.potion))
	c.PotionCount++

	time.Sleep(EndCraftDelay)

	c.startCraft()
}

func (c *Crafter) RunHooks() tea.Cmd {
	return func() tea.Msg {
		go func() {
			hook.Register(hook.KeyDown, []string{c.startPause}, func(e hook.Event) {
				if utils.Logger != nil {
					utils.Logger.Printf("Start/Pause key \"%s\" pressed\n", c.startPause)
				}

				if c.paused {
					c.StartProgram()
				} else {
					c.StopProgram()
				}
			})

			hook.Register(hook.KeyDown, []string{c.stop}, func(e hook.Event) {
				if utils.Logger != nil {
					utils.Logger.Printf("Stop key \"%s\" pressed\n", c.stop)
				}

				c.ExitProgram()
				hook.End()
			})

			s := hook.Start()

			if utils.Logger != nil {
				utils.Logger.Println("Initialize hooks")
			}

			for {
				select {
				case <-c.StopHookContext.Done():
					return
				default:
					<-hook.Process(s)
				}
			}
		}()

		return nil
	}
}

func (c *Crafter) GetCurrentAmount() int {
	return c.CurrentAmount
}
