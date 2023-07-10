package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Primary Style
	mainStyle = lipgloss.NewStyle().Margin(1, 1)

	// Colors
	Primary   = lipgloss.Color("#364F6B")
	Secondary = lipgloss.Color("#3FC1C9")
	Tertiary  = lipgloss.Color("#F5F5F5")
)

type Model struct {
	// List to show recipes
	List list.Model

	// Helpers
	width int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := mainStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
		m.width = msg.Width
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	listView := m.List.View()

	// Check if selectedItem is not nil before type assertion
	item := m.List.SelectedItem()
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

	// Apply mainStyle to listView and detailsView
	listView = mainStyle.Render(listView)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(m.width/2).Render(listView),
		detailsStyle,
	)
}
