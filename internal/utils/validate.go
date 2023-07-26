package utils

import (
	"errors"
	"fmt"
	"strings"
)

// Validate checks and validates the entire config for XIVCrafter
func Validate(startPause string, stop string, confirm string, cancel string, recipes []Recipe) (string, error) {
	// Validate all recipes
	for _, recipe := range recipes {
		if err := ValidateRecipe(startPause, stop, confirm, cancel, recipe); err != nil {
			return recipe.Name, err
		}
	}

	return "", nil
}

// ValidateSettings checks and validates just the XIVCrafter settings
func ValidateSettings(startPause string, stop string, confirm string, cancel string) error {
	keys := map[string]string{
		"StartPause": startPause,
		"Stop":       stop,
		"Confirm":    confirm,
		"Cancel":     cancel,
	}

	// Check if all hotkeys are unique per recipe
	if !CheckUniqueKeys(keys["StartPause"], keys["Stop"], keys["Confirm"], keys["Cancel"]) {
		return errors.New("hotkeys are not unique")
	}

	// Check if each hotkey is a valid key
	invalidKeys := []string{}
	for key, hotkey := range keys {
		if hotkey == "" || !CheckValidKey(hotkey) {
			invalidKeys = append(invalidKeys, key)
		}
	}
	if len(invalidKeys) > 0 {
		return fmt.Errorf("these are not valid keys: %v", invalidKeys)
	}

	return nil
}

// ValidateSettings checks and validates the XIVCrafter settings and a recipe
func ValidateRecipe(startPause string, stop string, confirm string, cancel string, recipe Recipe) error {
	keys := map[string]string{
		"StartPause": startPause,
		"Stop":       stop,
		"Confirm":    confirm,
		"Cancel":     cancel,
		"Food":       recipe.Food,
		"Potion":     recipe.Potion,
		"Macro1":     recipe.Macro1,
		"Macro2":     recipe.Macro2,
		"Macro3":     recipe.Macro3,
	}

	// Slice of keys that must not be an empty string
	requiredKeys := []string{"StartPause", "Stop", "Confirm", "Cancel", "Macro1"}

	// Check if all hotkeys are unique per recipe
	if !CheckUniqueKeys(keys["StartPause"], keys["Stop"], keys["Confirm"], keys["Cancel"], keys["Food"], keys["Potion"], keys["Macro1"], keys["Macro2"], keys["Macro3"]) {
		return errors.New("hotkeys are not unique")
	}

	// Check if each hotkey is a valid key
	invalidKeys := []string{}
	for key, hotkey := range keys {
		if !CheckValidKey(hotkey) {
			invalidKeys = append(invalidKeys, key)
		} else {
			// Check if the hotkey is a required key and if it is an empty string
			for _, req := range requiredKeys {
				if hotkey == "" && key == req {
					invalidKeys = append(invalidKeys, key)
				}
			}
		}
	}
	if len(invalidKeys) > 0 {
		return fmt.Errorf("these are not valid keys: %v", invalidKeys)
	}

	// Check if the food duration is valid
	switch recipe.FoodDuration {
	case 30, 40, 45:
		// Do nothing
	default:
		return errors.New("food duration is invalid, must be 30, 40, or 45")
	}

	// Check if the macro durations are valid
	macroDurations := map[string]int{
		"Macro1Duration": recipe.Macro1Duration,
		"Macro2Duration": recipe.Macro2Duration,
		"Macro3Duration": recipe.Macro3Duration,
	}
	invalidDurations := []string{}
	for name, duration := range macroDurations {
		if duration < 1 {
			invalidDurations = append(invalidDurations, name)
		}
	}
	if len(invalidKeys) > 0 {
		return fmt.Errorf("these have invalid durations: %v", invalidDurations)
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
