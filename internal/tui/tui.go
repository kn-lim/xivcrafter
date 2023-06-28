package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// type Model struct {
// }

// // NewModel returns a new model used for rendering the TUI
// func NewModel() *Model {
// 	return &Model{}
// }

type Model int
type tickMsg time.Time

var logo = "\n __  __ ___ __     __ ____               __  _\n\\ \\/ /|_ _|\\ \\   / // ___| _ __  __ _  / _|| |_  ___  _ __\n \\  /  | |  \\ \\ / /| |    | '__|/ _` || |_ | __|/ _ \\| '__|\n /  \\  | |   \\ V / | |___ | |  | (_| ||  _|| |_|  __/| |\n/_/\\_\\|___|   \\_/   \\____||_|   \\__,_||_|   \\__|\\___||_|\n"

func (m Model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		m--
		if m <= 0 {
			return m, tea.Quit
		}
		return m, tick()
	}

	return m, nil
}

func (m Model) View() string {
	message := fmt.Sprintf("Hi. This program will exit in %d seconds...", m)

	// Style message with a border
	messageStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Render(message)

	// Style the full screen with another border
	screenStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Align(lipgloss.Center).
		Render(messageStyle)

	return "\n" + screenStyle
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
