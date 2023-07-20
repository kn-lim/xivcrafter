package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

type List struct {
	// List to show recipes
	Recipes list.Model

	// Helpers
	width int
}

func (m List) Init() tea.Cmd {
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
			if utils.Logger != nil {
				utils.Logger.Printf("Selected Recipe Name: %s\n", m.Recipes.SelectedItem().(Item).Name)
			}

			Models[Recipes] = m
			return Models[Amount].Update(nil)
		}

	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
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
		selectedItem := item.(Item)
		detailsView = lipgloss.NewStyle().Render(selectedItem.PrintItemDetails())
		detailsView = listStyle.Padding(0, 2).Render(detailsView)

		detailsStyle = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(Secondary).
			MarginTop(6).
			Render(detailsView)
	}

	// Apply mainStyle to recipeView and detailsView
	recipeView = listStyle.Render(recipeView)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(m.width/2).Render(recipeView),
		detailsStyle,
	)
}
