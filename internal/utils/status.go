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
		lipgloss.NewStyle().Foreground(Gray).Render("Waiting"),
		lipgloss.NewStyle().Foreground(Green).Render("Crafting"),
		lipgloss.NewStyle().Foreground(Yellow).Render("Pausing"),
		lipgloss.NewStyle().Foreground(Yellow).Render("Paused"),
		lipgloss.NewStyle().Foreground(Red).Render("Stopping"),
		lipgloss.NewStyle().Foreground(Red).Render("Stopped"),
		lipgloss.NewStyle().Foreground(Green).Render("Finished"),
	}
)
