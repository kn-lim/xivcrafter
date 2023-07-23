package crafter

import "time"

var (
	// XIVCrafter delays
	PauseDelay = 1 * time.Second

	// FFXIV timings
	PotionDuration = 15 * time.Minute

	// Crafting delays
	KeyDelay        = 500 * time.Millisecond
	StartCraftDelay = 2 * time.Second
	EndCraftDelay   = 3 * time.Second
)
