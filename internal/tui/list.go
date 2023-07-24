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

	// XIVCrafter settings
	StartPause string
	Stop       string

	// In-game hotkeys
	Confirm string
	Cancel  string

	// Helpers
	width int
	msg   string
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

		// Add new recipe
		case "n":
			Models[Recipes] = m
			return Models[NewRecipe].Update(nil)

		// Edit recipe
		case "e":
			return m, nil
			// 	Models[Recipes] = m
			// 	return Models[EditRecipe].Update(nil)
		}

	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
		m.Recipes.SetSize(msg.Width-h, msg.Height-v)
		m.width = msg.Width

	case Item:
		if utils.Logger != nil {
			utils.Logger.Println("Updating list of recipes")
		}

		m.Recipes.SetItems(append(m.Recipes.Items(), msg))

		// Save to config
		itemsList := m.Recipes.Items()
		items := make([]Item, len(itemsList))
		for i, listItem := range itemsList {
			if item, ok := listItem.(Item); ok {
				items[i] = item
			}
		}
		utils.WriteToConfig(m.StartPause, m.Stop, m.Confirm, m.Cancel, convertItemsToRecipes(items))
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

func NewList(startPause string, stop string, confirm string, cancel string, items []list.Item) *List {
	model := &List{
		StartPause: startPause,
		Stop:       stop,
		Confirm:    confirm,
		Cancel:     cancel,
		Recipes:    list.New(items, NewItemDelegate(), 0, 0),
	}

	model.Recipes.Title = "XIVCrafter"
	model.Recipes.Styles.Title = model.Recipes.Styles.Title.Padding(1, 3, 1).Bold(true).Background(Primary).Foreground(Tertiary)

	return model
}
