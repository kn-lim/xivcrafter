package tui

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Recipe struct {
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

func (r Recipe) FilterValue() string { return r.Name }
func (r Recipe) Title() string       { return r.Name }
func (r Recipe) Description() string {
	// Get number of macros
	macroCount := 0
	if r.Macro1 != "" {
		macroCount++
	}
	if r.Macro2 != "" {
		macroCount++
	}
	if r.Macro3 != "" {
		macroCount++
	}

	// Get duration of macros in seconds
	recipeDuration := 0
	if r.Macro1 != "" {
		recipeDuration += r.Macro1Duration
	}
	if r.Macro2 != "" {
		recipeDuration += r.Macro2Duration
	}
	if r.Macro3 != "" {
		recipeDuration += r.Macro3Duration
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

// Helper function to print recipe details
func (r Recipe) PrintRecipeDetails() string {
	var details string
	boldStyle := lipgloss.NewStyle().Bold(true)

	details += boldStyle.Render("Name") + ": " + r.Name + "\n"

	if r.Food != "" {
		details += boldStyle.Render("Food") + ": " + r.Food + "\n"
	}

	if r.Potion != "" {
		details += boldStyle.Render("Potion") + ": " + r.Potion + "\n"
	}

	details += boldStyle.Render("Macro 1") + ": " + r.Macro1 + "\n"
	details += boldStyle.Render("Macro 1 Duration") + ": " + strconv.Itoa(r.Macro1Duration)

	if r.Macro2 != "" {
		details += "\n" + boldStyle.Render("Macro 2") + ": " + r.Macro2 + "\n"
		details += boldStyle.Render("Macro 2 Duration") + ": " + strconv.Itoa(r.Macro2Duration)
	}

	if r.Macro3 != "" {
		details += "\n" + boldStyle.Render("Macro 3") + ": " + r.Macro3 + "\n"
		details += boldStyle.Render("Macro 3 Duration") + ": " + strconv.Itoa(r.Macro3Duration)
	}

	return details
}

// NewRecipe returns the default settings for a Recipe struct
func NewRecipe() *Recipe {
	return &Recipe{
		Name:           "",
		Food:           "",
		FoodDuration:   30,
		Potion:         "",
		Macro1:         "",
		Macro1Duration: 1,
		Macro2:         "",
		Macro2Duration: 1,
		Macro3:         "",
		Macro3Duration: 1,
	}
}
