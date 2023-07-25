package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

type Item struct {
	// XIVCrafter Settings
	Name string

	// Consumables
	Food         string
	FoodDuration int
	Potion       string

	// In-game hotkeys
	Macro1         string
	Macro1Duration int
	Macro2         string
	Macro2Duration int
	Macro3         string
	Macro3Duration int
}

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

// convertItemsToRecipes takes a slice of Item structs and returns a slice of utils.Recipe structs
func convertItemsToRecipes(items []Item) []utils.Recipe {
	recipes := make([]utils.Recipe, len(items))

	for i, item := range items {
		recipes[i] = utils.Recipe{
			Name:           item.Name,
			Food:           item.Food,
			FoodDuration:   item.FoodDuration,
			Potion:         item.Potion,
			Macro1:         item.Macro1,
			Macro1Duration: item.Macro1Duration,
			Macro2:         item.Macro2,
			Macro2Duration: item.Macro2Duration,
			Macro3:         item.Macro3,
			Macro3Duration: item.Macro3Duration,
		}
	}

	return recipes
}

func updateItems(items []list.Item, newItem Item) []list.Item {
	found := false
	for i, item := range items {
		if item.(Item).Name == newItem.Name {
			if utils.Logger != nil {
				utils.Logger.Printf("Updating settings of recipe %s\n", newItem.Name)
			}

			// If the item exists, copy the settings of newItem into the existing item.
			items[i] = newItem
			found = true
			break
		}
	}

	if !found {
		if utils.Logger != nil {
			utils.Logger.Printf("Adding new recipe %s\n", newItem.Name)
		}

		// If the item doesn't exist, append newItem to the slice.
		items = append(items, newItem)
	}

	return items
}
