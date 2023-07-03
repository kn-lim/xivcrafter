package tui

import (
	"strconv"

	"github.com/kn-lim/xivcrafter/internal/utils"
)

type Item struct {
	Recipe utils.Recipe
}

func (i Item) Title() string { return i.Recipe.Name }
func (i Item) Description() string {
	// Get number of macros
	macroCount := 0
	if i.Recipe.Macro1 != "" {
		macroCount++
	}
	if i.Recipe.Macro2 != "" {
		macroCount++
	}
	if i.Recipe.Macro3 != "" {
		macroCount++
	}

	// Get duration of macros in seconds
	recipeDuration := 0
	if i.Recipe.Macro1 != "" {
		recipeDuration += i.Recipe.Macro1Duration
	}
	if i.Recipe.Macro2 != "" {
		recipeDuration += i.Recipe.Macro2Duration
	}
	if i.Recipe.Macro3 != "" {
		recipeDuration += i.Recipe.Macro3Duration
	}

	// Construct the description string with correct singular/plural forms
	macroStr := "Macros"
	if macroCount == 1 {
		macroStr = "Macro"
	}

	secondsStr := "Seconds"
	if recipeDuration == 1 {
		secondsStr = "Second"
	}

	return strconv.Itoa(macroCount) + " " + macroStr + ", " + strconv.Itoa(recipeDuration) + " " + secondsStr
}
func (i Item) FilterValue() string { return i.Recipe.Name }
