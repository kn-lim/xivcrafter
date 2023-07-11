package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

var tester bool = false

type tickMsg time.Time

type Progress struct {
	// Show crafting progress
	progress progress.Model
	percent  float64

	// Helpers
	msg string
}

func (m Progress) Init() tea.Cmd {
	return nil
}

func (m Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.msg = "updating"

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit XIVCrafter
		case "ctrl+c", "q":
			return m, tea.Quit

		// Go back to amount input model
		case "esc", "b":
			return models[AMOUNT].Update(nil)

		// Start progress bar
		case "enter":
			return m, tickCmd()
		}

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		m.percent += 0.25
		if m.percent > 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()

	default:
		return m, nil
	}

	return m, nil
}

func (m Progress) View() string {
	if !tester {
		m.msg = "test"
		tester = true
	}
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n" + m.msg
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
