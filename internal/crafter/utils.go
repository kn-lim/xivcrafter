package crafter

import "time"

const (
	// Delay whenever crafting is paused
	PauseDelay = 1 * time.Second

	// Duration of potion buff
	PotionDuration = 15 * time.Minute

	// Crafting delays

	// Delay after key press
	KeyDelay = 500 * time.Millisecond

	// Delay for starting craft animations
	StartCraftDelay = 1500 * time.Millisecond

	// Delay for ending craft animations
	EndCraftDelay = 2500 * time.Millisecond
)
