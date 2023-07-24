package utils

import (
	"errors"
	"fmt"
	"strings"
)

// Validate checks and validates the config file for XIVCrafter
func Validate(startPause string, stop string, confirm string, cancel string, recipe Recipe) error {
	keys := map[string]string{
		"start_pause": startPause,
		"stop":        stop,
		"confirm":     confirm,
		"cancel":      cancel,
		"food":        recipe.Food,
		"potion":      recipe.Potion,
		"macro1":      recipe.Macro1,
		"macro2":      recipe.Macro2,
		"macro3":      recipe.Macro3,
	}

	// Check if all hotkeys are unique
	if !CheckUniqueKeys(startPause, stop, confirm, cancel, recipe.Food, recipe.Potion, recipe.Macro1, recipe.Macro2, recipe.Macro3) {
		return errors.New("hotkeys are not unique")
	}

	// Check if the food duration is valid
	switch recipe.FoodDuration {
	case 30, 40, 45:
		// Do nothing
	default:
		return fmt.Errorf("%v is not valid. must be either 30, 40 or 45", recipe.FoodDuration)
	}

	// Check if each hotkey is a valid key
	for name, key := range keys {
		if !CheckValidKey(key) {
			return fmt.Errorf("%s is not a valid key", name)
		}
	}

	return nil
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
	alphaNumericKeys := []string{
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

	allKeys := append(alphaNumericKeys, functionKeys...)
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
