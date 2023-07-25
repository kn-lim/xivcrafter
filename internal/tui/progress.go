package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/crafter"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

const (
	padding     = 2
	maxWidth    = 100
	statusWidth = 10
)

var (
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
			// Exits currently running crafter
			m.Crafter.ExitProgram()

			// Reset crafter pointer to zero-valued Crafter struct
			m.Crafter = &crafter.Crafter{}

			return Models[Amount].Update(nil)
		}

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - padding*2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.Progress.Percent() == 1.0 {
			if utils.Logger != nil {
				utils.Logger.Println("Setting Status to \"Finished\"")
			}

			m.Crafter.Status = utils.Finished
			return m, nil
		}

		cmd := m.Progress.SetPercent(float64(m.Crafter.CurrentAmount) / float64(Models[Amount].(Input).amount))

		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	case initialize:
		if utils.Logger != nil {
			utils.Logger.Println("Initializing progress bar")
		}

		m.Crafter = crafter.NewCrafter(m.StartPause, m.Stop, m.Confirm, m.Cancel)

		recipe := Models[Recipes].(List).Recipes.SelectedItem().(Item)
		m.Crafter.SetRecipe(recipe.Name, Models[Amount].(Input).amount, recipe.Food, recipe.FoodDuration, recipe.Potion, recipe.Macro1, recipe.Macro1Duration, recipe.Macro2, recipe.Macro2Duration, recipe.Macro3, recipe.Macro3Duration)

		return m, tea.Batch(tickCmd(), m.Crafter.RunHooks(), m.Crafter.Run())

	default:
		return m, nil
	}

	return m, nil
}

func (m Progress) View() string {
	recipeView := lipgloss.NewStyle().Width(30).Render(Models[Recipes].(List).Recipes.SelectedItem().(Item).PrintItemDetails())
	hotkeyView := fmt.Sprintf("Press \"%s\" to Start/Pause\nPress \"%s\" to Stop", lipgloss.NewStyle().Bold(true).Render(m.StartPause), lipgloss.NewStyle().Bold(true).Render(m.Stop))
	configView := lipgloss.JoinHorizontal(lipgloss.Left, recipeView, hotkeyView)
	amountView := fmt.Sprintf("%s: %v", lipgloss.NewStyle().Bold(true).Render("\nAmount to Craft"), Models[Amount].(Input).amount)
	craftingView := lipgloss.NewStyle().PaddingLeft(3).Render(fmt.Sprintf("(%v/%v)", m.Crafter.CurrentAmount, Models[Amount].(Input).amount))
	progressView := lipgloss.JoinHorizontal(lipgloss.Left, lipgloss.NewStyle().Width(statusWidth).Render(utils.Status[m.Crafter.Status]), m.Progress.View(), craftingView)
	helpView := "\n\n\n" + m.Help.View(progressKeys)

	return mainStyle.Render(lipgloss.JoinVertical(lipgloss.Left, titleView, "", configView, amountView, "\n", progressView, helpView)) + "\n"
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewProgress(startPause string, stop string, confirm string, cancel string) *Progress {
	model := &Progress{
		Crafter:    &crafter.Crafter{},
		Progress:   progress.New(progress.WithGradient(string(ProgressStart), string(ProgressEnd))),
		Help:       help.New(),
		StartPause: startPause,
		Stop:       stop,
		Confirm:    confirm,
		Cancel:     cancel,
	}

	return model
}
