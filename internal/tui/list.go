package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type List struct {
	// List to show recipes
	Recipes list.Model

	// Helpers
	width int
}

func (m List) Init() tea.Cmd {
	// Setup Input model
	inputModel := Input{
		input:  textinput.New(),
		amount: 1,
	}
	inputModel.input.Focus()
	models[AMOUNT] = inputModel

	// Setup Progress model
	progressModel := Progress{
		progress: progress.New(),
		percent:  1,
	}
	models[CRAFTER] = progressModel

	return nil
}

func (m List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit XIVCrafter
		case "ctrl+c", "q":
			return m, tea.Quit

		// Select recipe to craft
		case "enter":
			models[RECIPES] = m
			return models[AMOUNT].Update(nil)
		}

	case tea.WindowSizeMsg:
		h, v := mainStyle.GetFrameSize()
		m.Recipes.SetSize(msg.Width-h, msg.Height-v)
		m.width = msg.Width
	}

	var cmd tea.Cmd
	m.Recipes, cmd = m.Recipes.Update(msg)
	return m, cmd
}

func (m List) View() string {
	recipeView := m.Recipes.View()

	// Check if selectedItem is not nil before type assertion
	item := m.Recipes.SelectedItem()
	var detailsView string
	var detailsStyle string
	if item != nil {
		selectedItem := item.(Recipe)
		detailsView = lipgloss.NewStyle().Foreground(Tertiary).Render(selectedItem.PrintRecipeDetails())
		detailsView = mainStyle.Padding(0, 2).Render(detailsView)

		detailsStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Secondary).
			MarginTop(6).
			Render(detailsView)
	}

	// Apply mainStyle to recipeView and detailsView
	recipeView = mainStyle.Render(recipeView)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(m.width/2).Render(recipeView),
		detailsStyle,
	)
}
