package utils

type Config struct {
	// XIVCrafter Hotkeys
	StartPause string `json:"start_pause"`
	Stop       string `json:"stop"`

	// In-Game Hotkeys
	Confirm string `json:"confirm"`
	Cancel  string `json:"cancel"`

	// Slice of Recipe structs
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	// XIVCrafter Settings
	Name     string `json:"name"`
	LastUsed bool   `json:"last_used"`
	Amount   int    `json:"amount"`

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

// NewConfig returns the default settings for a Config struct
func NewConfig() *Config {
	// Create variable to hold slice of recipes
	var recipes []Recipe
	recipes = append(recipes, *NewRecipe())

	// Set the default recipe to be marked as last used
	recipes[0].LastUsed = true

	return &Config{
		StartPause: "",
		Stop:       "",
		Confirm:    "",
		Cancel:     "",
		Recipes:    recipes,
	}
}

// NewRecipe returns the default settings for a Recipe struct
func NewRecipe() *Recipe {
	return &Recipe{
		Name:           "",
		LastUsed:       false,
		Amount:         1,
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
