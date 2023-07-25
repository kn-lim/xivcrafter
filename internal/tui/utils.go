package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

const (
	Recipes = iota
	UpdateRecipe
	Amount
	Crafter
)

const (
	MaxNameLength = 25
)

var (
	// Primary Style
	titleStyle = lipgloss.NewStyle().Padding(1, 3, 1).Bold(true).Background(utils.Primary).Foreground(utils.Default)
	titleView  = titleStyle.Render("XIVCrafter")
	listStyle  = lipgloss.NewStyle().Margin(1, 1)
	mainStyle  = lipgloss.NewStyle().Margin(1, 5)

	// Slice to manage models
	Models = []tea.Model{nil, nil, nil, nil, nil}
)
