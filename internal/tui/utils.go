package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Recipes = iota
	NewRecipe
	EditRecipe
	Amount
	Crafter
)

const (
	MaxNameLength = 25
)

var (
	// Primary Style
	titleView = lipgloss.NewStyle().MarginBottom(1).Padding(1, 3, 1).Bold(true).Background(Primary).Foreground(Tertiary).Render("XIVCrafter")
	listStyle = lipgloss.NewStyle().Margin(1, 1)
	mainStyle = lipgloss.NewStyle().Margin(1, 5)

	// Colors
	Primary    = lipgloss.Color("#364F6B")
	Secondary  = lipgloss.Color("#3FC1C9")
	Tertiary   = lipgloss.Color("#F5F5F5")
	Quaternary = lipgloss.Color("#FC5185")

	// Slice to manage models
	Models = []tea.Model{nil, nil, nil, nil, nil}
)
