package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 100
)

const (
	Waiting = iota
	Crafting
	Pausing
	Paused
	Stopping
	Stopped
)

var (
	status = []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("Waiting"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render("Crafting"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Pausing"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Paused"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Stopping"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Stopping"),
	}

	// Progress Bar Colors
	ProgressStart = lipgloss.Color("#1B6B93")
	ProgressEnd   = lipgloss.Color("#A2FF86")
)

type tickMsg time.Time

type Progress struct {
	// Show crafting progress
	Progress progress.Model

	// XIVCrafter Settings
	StartPause string
	Stop       string
	Confirm    string
	Cancel     string

	// Helpers
	Status int
	msg    string
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

		// Go back to amount input model
		case "esc", "b":
			return Models[Amount].Update(nil)

		// Start progress bar
		case "enter":
			m.Status = Crafting

			if m.Progress.Percent() == 1.0 {
				m.Progress.SetPercent(0)
				time.Sleep(2 * time.Second)
			}

			return m, tickCmd()
		}

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - padding*2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.Progress.Percent() == 1.0 {
			return m, nil
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.Progress.IncrPercent(0.20)
		return m, tea.Batch(tickCmd(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}

	return m, nil
}

func (m Progress) View() string {
	recipeView := lipgloss.NewStyle().Width(30).Render(Models[Recipes].(List).Recipes.SelectedItem().(Recipe).PrintRecipeDetails())
	hotkeyView := fmt.Sprintf("Press \"%s\" to Start/Pause\nPress \"%s\" to Stop", lipgloss.NewStyle().Bold(true).Render(m.StartPause), lipgloss.NewStyle().Bold(true).Render(m.Stop))
	configView := lipgloss.JoinHorizontal(lipgloss.Left, recipeView, hotkeyView)
	amountView := fmt.Sprintf("%s: %v", lipgloss.NewStyle().Bold(true).Render("\nAmount to Craft"), Models[Amount].(Input).amount)
	progressView := lipgloss.JoinHorizontal(lipgloss.Left, lipgloss.NewStyle().Width(10).Render(status[m.Status]), m.Progress.View())

	return mainStyle.Render(lipgloss.JoinVertical(lipgloss.Left, titleView, configView, amountView, "\n", progressView)) + "\n"
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
