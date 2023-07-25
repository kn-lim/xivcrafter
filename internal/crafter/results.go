package crafter

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

var (
	Results = []Result{}
)

type Result struct {
	Name        string
	CraftTime   time.Duration
	CraftCount  int
	FoodCount   int
	PotionCount int
}

// NewResults returns a pointer to a Result struct
func NewResult(name string, startTime time.Time, endTime time.Time, craftCount int, foodCount int, potionCount int) *Result {
	return &Result{
		Name:        name,
		CraftTime:   endTime.Sub(startTime),
		CraftCount:  craftCount,
		FoodCount:   foodCount,
		PotionCount: potionCount,
	}
}

// PrintResults prints the results to stdout once XIVCrafter stops running
func PrintResults() {
	var totalCraftTime time.Duration
	totalCraftCount := 0
	totalFoodCount := 0
	totalPotionCount := 0

	// Get totals
	for _, result := range Results {
		totalCraftTime += result.CraftTime
		totalCraftCount += result.CraftCount
		totalFoodCount += result.FoodCount
		totalPotionCount += result.PotionCount
	}

	output := fmt.Sprintf("\n%s\nTotal Crafting Time: %s\nTotal Amount Crafted: %v\n", lipgloss.NewStyle().MarginBottom(1).Padding(1, 3, 1).Bold(true).Background(utils.Primary).Foreground(utils.Default).Render("XIVCrafter Results"), totalCraftTime.Truncate(time.Second).String(), totalCraftCount)
	if totalFoodCount > 0 {
		output += fmt.Sprintf("Total Amount of Food Consumed: %v\n", totalFoodCount)
	}
	if totalPotionCount > 0 {
		output += fmt.Sprintf("Total Amount of Potions Consumed: %v\n", totalPotionCount)
	}
	fmt.Print(output + "\n")
}
