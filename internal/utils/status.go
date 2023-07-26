package utils

import "github.com/charmbracelet/lipgloss"

// Indices for status color codes
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
	// Status color codes
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
