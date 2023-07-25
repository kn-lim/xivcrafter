package utils

import (
	"errors"
	"fmt"
	"strings"
)

// Validate checks and validates the config file for XIVCrafter
func Validate(startPause string, stop string, confirm string, cancel string, recipes []Recipe) error {
	// Get all keys
	keys := []string{startPause, stop, confirm, cancel}
	for _, recipe := range recipes {
		keys = append(keys, recipe.Food, recipe.Potion, recipe.Macro1)

		if recipe.Macro2 != "" {
			keys = append(keys, recipe.Macro2)
		}

		if recipe.Macro3 != "" {
			keys = append(keys, recipe.Macro3)
		}
	}

	// Check if all recipe names are unique
	if !CheckUniqueNames(recipes) {
		return errors.New("recipe names are not unique")
	}

	for _, recipe := range recipes {
		// Check if all hotkeys are unique per recipe
		if !CheckUniqueKeys(startPause, stop, confirm, cancel, recipe.Food, recipe.Potion, recipe.Macro1, recipe.Macro2, recipe.Macro3) {
			return errors.New("hotkeys are not unique")
		}

		// Check if the recipe food duration is valid
		switch recipe.FoodDuration {
		case 30, 40, 45:
			// Do nothing
		default:
			return fmt.Errorf("%v is not valid. must be either 30, 40 or 45", recipe.FoodDuration)
		}
	}

	// Check if each hotkey is a valid key
	for _, key := range keys {
		if !CheckValidKey(key) {
			return fmt.Errorf("%s is not a valid key", key)
		}
	}

	return nil
}

// CheckUniqueNames checks to see if the Name field are unique per Recipe
func CheckUniqueNames(recipes []Recipe) bool {
	names := make(map[string]bool)

	for _, recipe := range recipes {
		if _, exists := names[recipe.Name]; exists {
			// Name is not unique
			return false
		}

		names[recipe.Name] = true
	}

	// All recipe names are unique
	return true
}

// CheckUniqueKeys checks to see if all hotkeys are unique
func CheckUniqueKeys(keys ...string) bool {
	keyMap := make(map[string]bool)

	for _, key := range keys {
		if key == "" {
			// Ignore unused hotkey
			continue
		}

		if _, exists := keyMap[key]; exists {
			// Hotkey is not unique
			return false
		}
		keyMap[key] = true
	}

	// All hotkeys are unique
	return true
}

// CheckValidKey checks to see if the given string is a valid hotkey for XIVCrafter
func CheckValidKey(key string) bool {
	alphanumericalKeys := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	functionKeys := []string{
		"f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8", "f9", "f10", "f11", "f12",
	}
	specialKeys := []string{
		"backspace", "delete", "enter", "tab", "escape", "space", "up", "down", "right", "left", "home", "end", "pageup", "pagedown",
	}
	modifierKeys := []string{
		"right_shift", "left_shift", "right_ctrl", "left_ctrl", "right_alt", "left_alt", "right_super", "left_super",
	}
	numberPadKeys := []string{
		"pad0", "pad1", "pad2", "pad3", "pad4", "pad5", "pad6", "pad7", "pad8", "pad9",
		"pad*", "pad+", "padenter", "pad.", "pad-", "pad/",
	}
	noKey := ""

	allKeys := append(alphanumericalKeys, functionKeys...)
	allKeys = append(allKeys, specialKeys...)
	allKeys = append(allKeys, modifierKeys...)
	allKeys = append(allKeys, numberPadKeys...)
	allKeys = append(allKeys, noKey)

	for _, validKey := range allKeys {
		if strings.EqualFold(key, validKey) {
			return true
		}
	}

	return false
}
