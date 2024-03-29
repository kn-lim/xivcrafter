package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

type Item struct {
	// XIVCrafter Settings

	// Name of recipe
	Name string

	// Consumables

	// Food hotkey
	Food string
	// Duration of food
	FoodDuration int
	// Potion hotkey
	Potion string

	// In-game hotkeys

	// Macro 1 hotkey
	Macro1 string
	// Duration of macro 1
	Macro1Duration int
	// Macro 2 hotkey
	Macro2 string
	// Duration of macro 2
	Macro2Duration int
	// Macro 3 hotkey
	Macro3 string
	// Duration of macro 3
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

// PrintItemDetails returns a stylized string with details of the Item struct
func (i Item) PrintItemDetails() string {
	var details string
	style := lipgloss.NewStyle().Foreground(utils.Secondary)

	details += style.Render("Name") + ": " + i.Name + "\n"

	if i.Food != "" {
		details += style.Render("Food") + ": " + i.Food + "\n"
	}

	if i.Potion != "" {
		details += style.Render("Potion") + ": " + i.Potion + "\n"
	}

	details += style.Render("Macro 1") + ": " + i.Macro1 + "\n"
	details += style.Render("Macro 1 Duration") + ": " + strconv.Itoa(i.Macro1Duration)

	if i.Macro2 != "" {
		details += "\n" + style.Render("Macro 2") + ": " + i.Macro2 + "\n"
		details += style.Render("Macro 2 Duration") + ": " + strconv.Itoa(i.Macro2Duration)
	}

	if i.Macro3 != "" {
		details += "\n" + style.Render("Macro 3") + ": " + i.Macro3 + "\n"
		details += style.Render("Macro 3 Duration") + ": " + strconv.Itoa(i.Macro3Duration)
	}

	return details
}

// ConvertItemsToRecipes takes a slice of Item structs and returns a slice of utils.Recipe structs
func ConvertItemsToRecipes(items []Item) []utils.Recipe {
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

// ConvertListItemToItem takes in a list.Item slice and converts it into an Item slice
func ConvertListItemToItem(listItems []list.Item) []Item {
	items := make([]Item, len(listItems))
	for i, listItem := range listItems {
		if item, ok := listItem.(Item); ok {
			items[i] = item
		}
	}
	return items
}
