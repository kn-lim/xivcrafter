package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/crafter"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

const (
	// Progress bar visual settings

	padding     = 2
	maxWidth    = 100
	statusWidth = 10

	// Starting color of the progress bar
	ProgressStart = lipgloss.Color("#1B6B93")
	// Ending color of the progress bar
	ProgressEnd = lipgloss.Color("#A2FF86")
)

type Progress struct {
	// Show crafting progress
	Crafter  *crafter.Crafter
	Progress progress.Model

	// Help model
	Help help.Model

	// Helpers
	amount int
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
		WindowWidth = msg.Width
		WindowHeight = msg.Height

		m.Progress.Width = msg.Width - padding*2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil

	// From tickCmd
	case tickMsg:
		if m.Progress.Percent() == 1.0 {
			return m, nil
		}

		cmd := m.Progress.SetPercent(float64(m.Crafter.CurrentAmount) / float64(m.amount))

		return m, tea.Batch(tickCmd(), cmd)

	// Progress bar animation
	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	// From Input model
	case int:
		// Save to Progress model
		m.amount = msg

		if utils.Logger != nil {
			utils.Logger.Println("Initializing progress bar")
		}

		// Start crafter
		m.Crafter = crafter.NewCrafter(StartPause, Stop, Confirm, Cancel)
		recipe := Models[Recipes].(List).Recipes.SelectedItem().(Item)
		m.Crafter.SetRecipe(recipe.Name, msg, recipe.Food, recipe.FoodDuration, recipe.Potion, recipe.Macro1, recipe.Macro1Duration, recipe.Macro2, recipe.Macro2Duration, recipe.Macro3, recipe.Macro3Duration)

		// Start goroutines
		return m, tea.Batch(tickCmd(), m.Crafter.RunHooks(), m.Crafter.Run())

	default:
		return m, nil
	}

	return m, nil
}

func (m Progress) View() string {
	_, v := utils.MainStyle.GetFrameSize()

	recipeView := lipgloss.NewStyle().Width(30).Render(Models[Recipes].(List).Recipes.SelectedItem().(Item).PrintItemDetails())
	hotkeyView := fmt.Sprintf("Press \"%s\" to Start/Pause\nPress \"%s\" to Stop", lipgloss.NewStyle().Foreground(utils.Tertiary).Render(StartPause), lipgloss.NewStyle().Foreground(utils.Tertiary).Render(Stop))
	amountView := fmt.Sprintf("%s: %v", lipgloss.NewStyle().Foreground(utils.Secondary).Render("Amount to Craft"), m.amount)
	craftingView := lipgloss.NewStyle().PaddingLeft(3).Render(fmt.Sprintf("(%v/%v)", m.Crafter.CurrentAmount, m.amount))

	mainView := lipgloss.JoinVertical(
		lipgloss.Top,
		utils.TitleView,
		"",
		lipgloss.JoinHorizontal(lipgloss.Left, recipeView, hotkeyView),
		"",
		amountView,
		"",
		lipgloss.JoinHorizontal(lipgloss.Left, lipgloss.NewStyle().Width(statusWidth).Render(utils.Status[m.Crafter.Status]), m.Progress.View(), craftingView),
	)
	mainView = lipgloss.NewStyle().Height(WindowHeight - v - 1).Render(mainView)

	return utils.MainStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		mainView,
		m.Help.View(progressKeys),
	))
}

// NewProgress returns a pointer to a Progress struct
func NewProgress() *Progress {
	model := &Progress{
		Crafter:  &crafter.Crafter{},
		Progress: progress.New(progress.WithGradient(string(ProgressStart), string(ProgressEnd))),
		Help:     help.New(),
	}

	return model
}
