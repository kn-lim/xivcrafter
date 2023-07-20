package tui

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	// XIVCrafter Settings
	Name string `json:"name"`

	// Consumables
	Food         string `json:"food"`
	FoodDuration int    `json:"food_duration"`
	Potion       string `json:"potion"`

	// In-Game Hotkeys
	Macro1         string `json:"macro1"`
	Macro1Duration int    `json:"macro1_duration"`
	Macro2         string `json:"macro2"`
	Macro2Duration int    `json:"macro2_duration"`
	Macro3         string `json:"macro3"`
	Macro3Duration int    `json:"macro3_duration"`
}

// Implement Item interface

func (i Item) FilterValue() string { return i.Name }
func (i Item) Title() string       { return i.Name }
func (i Item) Description() string {
	// Get number of macros
	macroCount := 0
	if i.Macro1 != "" {
		macroCount++
	}
	if i.Macro2 != "" {
		macroCount++
	}
	if i.Macro3 != "" {
		macroCount++
	}

	// Get duration of macros in seconds
	recipeDuration := 0
	if i.Macro1 != "" {
		recipeDuration += i.Macro1Duration
	}
	if i.Macro2 != "" {
		recipeDuration += i.Macro2Duration
	}
	if i.Macro3 != "" {
		recipeDuration += i.Macro3Duration
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

// Helper function to print item details
func (i Item) PrintItemDetails() string {
	var details string
	boldStyle := lipgloss.NewStyle().Bold(true)

	details += boldStyle.Render("Name") + ": " + i.Name + "\n"

	if i.Food != "" {
		details += boldStyle.Render("Food") + ": " + i.Food + "\n"
	}

	if i.Potion != "" {
		details += boldStyle.Render("Potion") + ": " + i.Potion + "\n"
	}

	details += boldStyle.Render("Macro 1") + ": " + i.Macro1 + "\n"
	details += boldStyle.Render("Macro 1 Duration") + ": " + strconv.Itoa(i.Macro1Duration)

	if i.Macro2 != "" {
		details += "\n" + boldStyle.Render("Macro 2") + ": " + i.Macro2 + "\n"
		details += boldStyle.Render("Macro 2 Duration") + ": " + strconv.Itoa(i.Macro2Duration)
	}

	if i.Macro3 != "" {
		details += "\n" + boldStyle.Render("Macro 3") + ": " + i.Macro3 + "\n"
		details += boldStyle.Render("Macro 3 Duration") + ": " + strconv.Itoa(i.Macro3Duration)
	}

	return details
}
