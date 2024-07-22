package crafter

import (
	"context"
	"fmt"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/spf13/cobra"

	"github.com/kn-lim/xivcrafter/internal/utils"
)

type Crafter struct {
	// XIVCrafter settings

	name       string
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

	Status          int
	startTime       time.Time
	endTime         time.Time
	CurrentAmount   int
	FoodCount       int
	foodStartTime   time.Time
	PotionCount     int
	potionStartTime time.Time

	// Helpers

	running               bool
	Paused                bool
	exitOnce              sync.Once
	StopCrafterContext    context.Context
	StopCrafterCancelFunc context.CancelFunc
	StopHooksContext      context.Context
	StopHooksCancelFunc   context.CancelFunc
}

// NewCrafter returns a pointer to a Crafter struct
func NewCrafter(startPause string, stop string, confirm string, cancel string) *Crafter {
	// Create contexts
	crafterCtx, crafterCancelFunc := context.WithCancel(context.Background())
	hookCtx, hookCancelFunc := context.WithCancel(context.Background())

	return &Crafter{
		// XIVCrafter settings
		// name: "",
		amount:     1,
		startPause: startPause,
		stop:       stop,

		// Consumables
		// food: "",
		// potion: "",

		// In-game hotkeys
		confirm: confirm,
		cancel:  cancel,
		// macro1: "",
		macro1Duration: 1 * time.Second,
		// macro2: "",
		macro2Duration: 1 * time.Second,
		// macro3: "",
		macro3Duration: 1 * time.Second,

		// Information
		Status:    utils.Waiting,
		startTime: time.Time{},
		endTime:   time.Time{},
		// CurrentAmount: 0,
		// FoodCount: 0,
		foodStartTime: time.Time{},
		// PotionCount: 0,
		potionStartTime: time.Time{},

		// Helpers
		running:               true,
		Paused:                true,
		StopCrafterContext:    crafterCtx,
		StopCrafterCancelFunc: crafterCancelFunc,
		StopHooksContext:      hookCtx,
		StopHooksCancelFunc:   hookCancelFunc,
	}
}

func (c *Crafter) SetRecipe(name string, amount int, food string, foodDuration int, potion string, macro1 string, macro1Duration int, macro2 string, macro2Duration int, macro3 string, macro3Duration int) {
	// XIVCrafter settings
	c.name = name
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

	utils.Log("Infow", "Using recipe",
		"name", name,
		"amount", amount,
		"food", food,
		"foodDuration", foodDuration,
		"potion", potion,
		"macro1", macro1,
		"macro1Duration", macro1Duration,
		"macro2", macro2,
		"macro2Duration", macro2Duration,
		"macro3", macro3,
		"macro3Duration", macro3Duration,
	)
}

// Run provides the main logic to handle crafting
func (c *Crafter) Run() tea.Cmd {
	return func() tea.Msg {
		go func() {
			utils.Log("Infow", "initializing crafter")

			// Loop if XIVCrafter is running
			for c.running {
				select {
				case <-c.StopCrafterContext.Done():
					// Exit
					utils.Log("Infow", "crafter context stopped")
					utils.Log("Infow", "setting status to \"Stopped\"")

					c.Status = utils.Stopped
					return
				default:
					// Loop if XIVCrafter is crafting
					for !c.Paused {
						utils.Log("Infow", "starting craft",
							"status", fmt.Sprintf("%v / %v", c.CurrentAmount+1, c.amount),
						)

						c.Status = utils.Crafting

						// Get the start crafting time
						if c.startTime.IsZero() {
							time.Sleep(PauseDelay)
							c.startTime = time.Now()
						}

						c.startCraft()

						// Check food
						select {
						case <-c.StopCrafterContext.Done():
							// Exit
							utils.Log("Infow", "crafter context stopped")
							utils.Log("Infow", "setting status to \"Stopped\"")

							c.Status = utils.Stopped
							return
						default:
							if c.food != "" {
								c.checkFood()
							}
						}

						// Check potion
						select {
						case <-c.StopCrafterContext.Done():
							// Exit
							utils.Log("Infow", "crafter context stopped")
							utils.Log("Infow", "setting status to \"Stopped\"")

							c.Status = utils.Stopped
							return
						default:
							if c.potion != "" {
								c.checkPotion()
							}
						}

						// Macro 1
						select {
						case <-c.StopCrafterContext.Done():
							// Exit
							utils.Log("Infow", "crafter context stopped")
							utils.Log("Infow", "setting status to \"Stopped\"")

							c.Status = utils.Stopped
							return
						default:
							// Activate macro 1
							utils.Log("Infow", "Activating macro 1")

							cobra.CheckErr(robotgo.KeyTap(c.macro1))
							time.Sleep(KeyDelay)
							time.Sleep(c.macro1Duration)
						}

						// Macro 2
						if c.macro2 != "" {
							select {
							case <-c.StopCrafterContext.Done():
								// Exit
								utils.Log("Infow", "crafter context stopped")
								utils.Log("Infow", "setting status to \"Stopped\"")

								c.Status = utils.Stopped
								return
							default:
								// Activate macro 2
								utils.Log("Infow", "activating macro 2")

								cobra.CheckErr(robotgo.KeyTap(c.macro2))
								time.Sleep(KeyDelay)
								time.Sleep(c.macro2Duration)
							}
						}

						// Macro 3
						if c.macro3 != "" {
							select {
							case <-c.StopCrafterContext.Done():
								// Exit
								utils.Log("Infow", "crafter context stopped")
								utils.Log("Infow", "setting status to \"Stopped\"")

								c.Status = utils.Stopped
								return
							default:
								// Activate macro 3
								utils.Log("Infow", "activating macro 3")

								cobra.CheckErr(robotgo.KeyTap(c.macro3))
								time.Sleep(KeyDelay)
								time.Sleep(c.macro3Duration)
							}
						}

						c.CurrentAmount++

						utils.Log("Infow", "finished craft",
							"status", fmt.Sprintf("%v / %v", c.CurrentAmount, c.amount),
						)

						// Exit if finished crafting
						if c.CurrentAmount >= c.amount {
							utils.Log("Infow", "setting status to \"Finished\"")

							c.Status = utils.Finished
							c.ExitProgram()
						}

						time.Sleep(EndCraftDelay)
					}
				}

				// Check to see if XIVCrafter is paused
				if c.Paused && c.running && !c.startTime.IsZero() && c.CurrentAmount < c.amount && c.Status != utils.Paused {
					utils.Log("Infow", "setting status to \"Paused\"")

					c.Status = utils.Paused
				}

				time.Sleep(PauseDelay)
			}

			// Exit XIVCrafter before finish
			if c.CurrentAmount < c.amount {
				utils.Log("Infow", "setting status to \"Stopped\"")

				c.Status = utils.Stopped
			}
		}()

		return nil
	}
}

// StartProgram sets the paused value to false
func (c *Crafter) StartProgram() {
	utils.Log("Infow", "starting xivcrafter")
	utils.Log("Infow", "setting status to \"Crafting\"")

	c.Paused = false
	c.Status = utils.Crafting
}

// StopProgram sets the paused value to true
func (c *Crafter) StopProgram() {
	utils.Log("Infow", "pausing xivcrafter")
	utils.Log("Infow", "setting status to \"Pausing\"")

	c.Paused = true
	c.Status = utils.Pausing
}

// ExitProgram sets the running value to false and the paused value to true
func (c *Crafter) ExitProgram() {
	c.exitOnce.Do(func() {
		utils.Log("Infow", "exiting xivcrafter")

		c.running = false
		c.Paused = true

		if c.CurrentAmount < c.amount {
			utils.Log("Infow", "setting status to \"Stopping\"")

			c.Status = utils.Stopping
		}

		utils.Log("Infow", "updating crafting results")

		Results = append(Results, *NewResult(c.name, c.startTime, time.Now(), c.CurrentAmount, c.FoodCount, c.PotionCount))

		utils.Log("Infow", "running the cancel functions for crafter and hooks")

		c.StopCrafterCancelFunc()
		c.StopHooksCancelFunc()
	})
}

// startCraft sets up the crafting action
func (c *Crafter) startCraft() {
	utils.Log("Infow", "starting crafting action")

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
	utils.Log("Infow", "stopping crafting action")

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
	utils.Log("Infow", "checking food")

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
	utils.Log("Infow", "consuming food",
		"count", c.FoodCount+1,
	)

	c.stopCraft()

	c.foodStartTime = time.Now()
	cobra.CheckErr(robotgo.KeyTap(c.food))
	c.FoodCount++

	time.Sleep(EndCraftDelay)

	c.startCraft()
}

// checkPotion checks to see whether the potion buff needs to be renewed
func (c *Crafter) checkPotion() {
	utils.Log("Infow", "checking potion")

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
	utils.Log("Infow", "consuming potion",
		"count", c.PotionCount+1,
	)

	c.stopCraft()

	c.potionStartTime = time.Now()
	cobra.CheckErr(robotgo.KeyTap(c.potion))
	c.PotionCount++

	time.Sleep(EndCraftDelay)

	c.startCraft()
}

// RunHooks provides the main logic to handle keyboard hook events
func (c *Crafter) RunHooks() tea.Cmd {
	return func() tea.Msg {
		go func() {
			hook.Register(hook.KeyDown, []string{c.startPause}, func(e hook.Event) {
				utils.Log("Infow", "start/pause key pressed",
					"key", c.startPause,
				)

				if c.Paused {
					c.StartProgram()
				} else {
					c.StopProgram()
				}
			})

			hook.Register(hook.KeyDown, []string{c.stop}, func(e hook.Event) {
				utils.Log("Infow", "stop key pressed",
					"key", c.stop,
				)

				c.ExitProgram()
			})

			s := hook.Start()

			utils.Log("Infow", "initializing hooks")

			for {
				select {
				case <-c.StopHooksContext.Done():
					utils.Log("Infow", "hook context stopped")

					hook.End()
					return
				case <-hook.Process(s):
					// Do nothing
				}
			}
		}()

		return nil
	}
}
