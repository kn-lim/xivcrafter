package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/crafter"
	hook "github.com/robotn/gohook"
)

const (
	padding     = 2
	maxWidth    = 100
	statusWidth = 10
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
	// Crafter Status Color Codes
	status = []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("Waiting"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render("Crafting"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Pausing"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Render("Paused"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Stopping"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Stopped"),
	}

	// Progress Bar Colors
	ProgressStart = lipgloss.Color("#1B6B93")
	ProgressEnd   = lipgloss.Color("#A2FF86")
)

type tickMsg time.Time

// Tells progress to initialize the crafter
type initialize struct{}

type Progress struct {
	// Show crafting progress
	Crafter  *crafter.Crafter
	Progress progress.Model

	// Help component
	Help help.Model

	// XIVCrafter settings
	StartPause string
	Stop       string

	// In-game hotkeys
	Confirm string
	Cancel  string

	// Helpers
	Status int
	// msg    string
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
			return Models[Amount].Update(backFromProgress{})
		}

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - padding*2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.Progress.Percent() == 1.0 {
			m.Status = Stopped
			return m, nil
		}

		cmd := m.Progress.SetPercent(float64(m.Crafter.CurrentAmount) / float64(Models[Amount].(Input).amount))

		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	case initialize:
		recipe := Models[Recipes].(List).Recipes.SelectedItem().(Recipe)
		m.Crafter.SetRecipe(Models[Amount].(Input).amount, recipe.Food, recipe.FoodDuration, recipe.Potion, recipe.Macro1, recipe.Macro1Duration, recipe.Macro2, recipe.Macro2Duration, recipe.Macro3, recipe.Macro3Duration)

		return m, tea.Batch(tickCmd(), m.startHooks())

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
	progressView := lipgloss.JoinHorizontal(lipgloss.Left, lipgloss.NewStyle().Width(statusWidth).Render(status[m.Status]), m.Progress.View())
	helpView := "\n\n\n" + m.Help.View(progressKeys)

	return mainStyle.Render(lipgloss.JoinVertical(lipgloss.Left, titleView, configView, amountView, "\n", progressView, helpView)) + "\n"
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Progress) startHooks() tea.Cmd {
	return func() tea.Msg {
		go func() {
			hook.Register(hook.KeyDown, []string{m.StartPause}, func(e hook.Event) {
				if m.Crafter.Paused {
					m.Status = Crafting
					m.Crafter.StartProgram()
					m.Update(nil)
				} else {
					m.Status = Paused
					m.Crafter.StopProgram()
					m.Update(nil)
				}
			})

			hook.Register(hook.KeyDown, []string{m.Stop}, func(e hook.Event) {
				m.Status = Stopping
				m.Crafter.ExitProgram()
				m.Update(nil)
				hook.End()
			})

			s := hook.Start()
			<-hook.Process(s)
		}()

		return nil
	}
}
