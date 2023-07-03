package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	List  list.Model
	Width int
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
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
		m.Width = msg.Width - h
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	listView := m.List.View()
	selectedItem := m.List.SelectedItem().(Item)
	detailsView := fmt.Sprintf("%s\n%s", selectedItem.Title(), selectedItem.Description())

	// Apply docStyle to listView and detailsView for margins
	listView = docStyle.Render(listView)
	detailsView = docStyle.Render(detailsView)

	// Position the detailsView to the right of the listView
	return lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(m.Width/2).Render(listView),
		lipgloss.NewStyle().Width(m.Width/2).Padding(4).Render(detailsView),
	)
}
