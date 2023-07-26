package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// For tick function
type tickMsg time.Time

// Indices for Models slice
const (
	Recipes = iota
	ChangeSettings
	ChangeRecipe
	Amount
	Crafter
)

var (
	// Height of terminal window
	WindowHeight int
	// Width of terminal window
	WindowWidth int

	// Slice to manage models
	Models = []tea.Model{nil, nil, nil, nil, nil}

	// Hotkeys

	StartPause string
	Stop       string
	Confirm    string
	Cancel     string
)

// Tick function
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
