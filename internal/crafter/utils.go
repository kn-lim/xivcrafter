package crafter

import "time"

var (
	// XIVCrafter delays
	PauseDelay = 1 * time.Second

	// FFXIV timings
	PotionDuration = 15 * time.Minute

	// Crafting delays
	KeyDelay        = 500 * time.Millisecond
	StartCraftDelay = 1500 * time.Millisecond
	EndCraftDelay   = 2500 * time.Millisecond
)
