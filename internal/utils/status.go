package utils

import "github.com/charmbracelet/lipgloss"

const (
	Waiting = iota
	Crafting
	Pausing
	Paused
	Stopping
	Stopped
	Finished
)

var (
	// Crafter Status Color Codes
	Status = []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("Waiting"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render("Crafting"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Pausing"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Paused"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Stopping"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Stopped"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render("Finished"),
	}
)
