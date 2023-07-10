package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Progress struct {
	// Show crafting progress
	progress progress.Model
	percent  float64
}

func (m Progress) Init() tea.Cmd {
	return nil
}

func (m Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit XIVCrafter
		case "ctrl+c", "q":
			return m, tea.Quit

		// Go back to recipes
		case "esc", "b":
			return models[AMOUNT].Update(nil)
		}
	}

	return m, nil
}

func (m Progress) View() string {
	return fmt.Sprintf("Amount: %s", strconv.Itoa(models[AMOUNT].(Input).amount))
}
