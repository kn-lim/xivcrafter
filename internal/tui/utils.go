package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	RECIPES status = iota
	AMOUNT
	CRAFTER
)

var (
	// Primary Style
	mainStyle = lipgloss.NewStyle().Margin(1, 1)

	// Colors
	Primary    = lipgloss.Color("#364F6B")
	Secondary  = lipgloss.Color("#3FC1C9")
	Tertiary   = lipgloss.Color("#F5F5F5")
	Quaternary = lipgloss.Color("#FC5185")

	// Slice to manage models
	models = []tea.Model{nil, nil, nil}
)
