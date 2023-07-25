package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

const (
	Recipes = iota
	ChangeSettings
	ChangeRecipe
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

	// Hotkeys
	StartPause string
	Stop       string
	Confirm    string
	Cancel     string
)

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
