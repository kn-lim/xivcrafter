package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
	"github.com/spf13/cobra"
)

type List struct {
	// List to show recipes
	Recipes list.Model

	// Helpers
	width int
	msg   string
}

func (m List) Init() tea.Cmd {
	return tickCmd()
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

		// Change XIVCrafter settings
		case "s":
			Models[Recipes] = m
			return Models[ChangeSettings].Update(nil)

		// Add new recipe
		case "n":
			Models[Recipes] = m
			return Models[ChangeRecipe].Update(nil)

		// Edit recipe
		case "e":
			Models[Recipes] = m
			return Models[ChangeRecipe].Update(m.Recipes.SelectedItem().(Item))
		}

	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
		m.Recipes.SetSize(msg.Width-h, msg.Height-v)
		m.width = msg.Width

	case Settings:
		if utils.Logger != nil {
			utils.Logger.Println("Updating XIVCrafter settings")
		}

		// Update hotkeys
		StartPause = msg.startPause
		Stop = msg.stop
		Confirm = msg.confirm
		Cancel = msg.cancel

		// Save to config
		itemsList := m.Recipes.Items()
		var items []Item
		if len(itemsList) == 0 {
			// Create new recipe
			return Models[ChangeRecipe].Update(nil)
		} else {
			items = make([]Item, len(itemsList))
			for i, listItem := range itemsList {
				if item, ok := listItem.(Item); ok {
					items[i] = item
				}
			}
		}
		if err := utils.WriteToConfig(StartPause, Stop, Confirm, Cancel, convertItemsToRecipes(items)); err != nil {
			cobra.CheckErr(err)
		}

		m.msg = lipgloss.NewStyle().Padding(0, 0, 0, 4).Bold(true).Foreground(utils.Green).Render("Saved XIVCrafter settings")
		return m, m.Recipes.NewStatusMessage(m.msg)

	case Item:
		if utils.Logger != nil {
			utils.Logger.Println("Updating list of recipes")
		}

		m.Recipes.SetItems(updateItems(m.Recipes.Items(), msg))

		// Save to config
		itemsList := m.Recipes.Items()
		items := make([]Item, len(itemsList))
		for i, listItem := range itemsList {
			if item, ok := listItem.(Item); ok {
				items[i] = item
			}
		}
		if err := utils.WriteToConfig(StartPause, Stop, Confirm, Cancel, convertItemsToRecipes(items)); err != nil {
			cobra.CheckErr(err)
		}

		m.msg = lipgloss.NewStyle().Padding(0, 0, 0, 4).Bold(true).Foreground(utils.Green).Render(fmt.Sprintf("Saved %s recipe", msg.Name))
		return m, m.Recipes.NewStatusMessage(m.msg)
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
			BorderForeground(utils.Secondary).
			MarginTop(6).
			Render(detailsView)
	}

	// Apply listStyle to recipeView
	recipeView = listStyle.Render(recipeView)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(m.width/2).Render(recipeView),
		detailsStyle,
	)
}

func NewList(items []list.Item) *List {
	model := &List{
		Recipes: list.New(items, NewItemDelegate(), 0, 0),
	}

	// Defaults
	model.Recipes.Title = "XIVCrafter"
	model.Recipes.Styles.Title = titleStyle
	model.Recipes.StatusMessageLifetime = 5 * time.Second

	// Additional help keys
	model.Recipes.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.enter,
			listKeys.newRecipe,
			listKeys.editRecipe,
		}
	}
	model.Recipes.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.enter,
			listKeys.changeSettings,
			listKeys.newRecipe,
			listKeys.editRecipe,
		}
	}

	return model
}
