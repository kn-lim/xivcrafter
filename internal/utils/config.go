package utils

import "github.com/kn-lim/xivcrafter/internal/tui"

type Config struct {
	// XIVCrafter Hotkeys
	StartPause string `json:"start_pause"`
	Stop       string `json:"stop"`

	// In-Game Hotkeys
	Confirm string `json:"confirm"`
	Cancel  string `json:"cancel"`

	// Slice of Recipe structs
	Recipes []tui.Recipe `json:"recipes"`
}

// NewConfig returns the default settings for a Config struct
func NewConfig() *Config {
	var recipes []tui.Recipe
	recipes = append(recipes, *tui.NewRecipe())

	return &Config{
		StartPause: "",
		Stop:       "",
		Confirm:    "",
		Cancel:     "",
		Recipes:    recipes,
	}
}
