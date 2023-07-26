package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kn-lim/xivcrafter/internal/utils"
	"github.com/spf13/cobra"
)

const StatusMsgLifetime = 10 * time.Second

type List struct {
	// List model
	Recipes list.Model

	// Status message
	msg string
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

		// Select recipe to delete
		case "x":
			recipeName := m.Recipes.SelectedItem().(Item).Name
			if utils.Logger != nil {
				utils.Logger.Printf("Deleting recipe: %s\n", recipeName)
			}

			// Delete item from List model
			index := m.Recipes.Index()
			m.Recipes.RemoveItem(index)
			if index >= len(m.Recipes.Items()) {
				m.Recipes.Select(len(m.Recipes.Items()) - 1)
			} else {
				m.Recipes.Select(index)
			}

			// Save to config
			items := ConvertListItemToItem(m.Recipes.Items())
			if err := utils.WriteToConfig(StartPause, Stop, Confirm, Cancel, ConvertItemsToRecipes(items)); err != nil {
				cobra.CheckErr(err)
			}

			m.msg = utils.ListStatusStyle.Foreground(utils.Green).Render(fmt.Sprintf("Deleted recipe: %s", recipeName))
			return m, m.Recipes.NewStatusMessage(m.msg)

		// Select recipe to craft
		case "enter":
			if utils.Logger != nil {
				utils.Logger.Printf("Selected recipe: %s\n", m.Recipes.SelectedItem().(Item).Name)
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
		h, v := utils.ListStyle.GetFrameSize()
		m.Recipes.SetSize(msg.Width-h, msg.Height-v)

	// From UpdateSettings model
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
		listItems := m.Recipes.Items()
		var items []Item
		if len(listItems) == 0 {
			// Create new recipe
			return Models[ChangeRecipe].Update(nil)
		} else {
			items = ConvertListItemToItem(listItems)
		}
		if err := utils.WriteToConfig(StartPause, Stop, Confirm, Cancel, ConvertItemsToRecipes(items)); err != nil {
			cobra.CheckErr(err)
		}

		m.msg = utils.ListStatusStyle.Foreground(utils.Green).Render("Saved XIVCrafter settings")
		return m, m.Recipes.NewStatusMessage(m.msg)

	// From UpdateRecipe model
	case Item:
		if utils.Logger != nil {
			utils.Logger.Println("Updating list of recipes")
		}

		m.Recipes.SetItems(updateItems(m.Recipes.Items(), msg))

		// Save to config
		items := ConvertListItemToItem(m.Recipes.Items())
		if err := utils.WriteToConfig(StartPause, Stop, Confirm, Cancel, ConvertItemsToRecipes(items)); err != nil {
			cobra.CheckErr(err)
		}

		m.msg = utils.ListStatusStyle.Foreground(utils.Green).Render(fmt.Sprintf("Saved recipe: %s", msg.Name))
		return m, m.Recipes.NewStatusMessage(m.msg)
	}

	var cmd tea.Cmd
	m.Recipes, cmd = m.Recipes.Update(msg)
	return m, cmd
}

func (m List) View() string {
	return utils.ListStyle.Render(m.Recipes.View())
}

// NewList returns a pointer to a List struct
func NewList(items []list.Item) *List {
	model := &List{
		Recipes: list.New(items, NewItemDelegate(), 0, 0),
	}

	// Defaults
	model.Recipes.Title = "XIVCrafter"
	model.Recipes.Styles.Title = utils.TitleStyle
	model.Recipes.StatusMessageLifetime = StatusMsgLifetime
	model.Recipes.SetFilteringEnabled(false) // disabled for now

	// Additional help keys
	model.Recipes.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.enter,
			listKeys.new,
			listKeys.edit,
			listKeys.delete,
		}
	}
	model.Recipes.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.enter,
			listKeys.change,
			listKeys.new,
			listKeys.edit,
			listKeys.delete,
		}
	}

	return model
}
