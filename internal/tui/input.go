package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Input struct {
	// Get user input for the amount to craft
	Input  textinput.Model
	amount int

	// Helpers
	msg string
}

func (m Input) Init() tea.Cmd {
	return nil
}

func (m Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit XIVCrafter
		case "ctrl+c", "q":
			return m, tea.Quit

		// Go back to recipes model
		case "esc", "b":
			return Models[Recipes].Update(nil)

		case "enter":
			// Check if the input model is focused
			if m.Input.Focused() {
				// Check if user input is a valid integer
				var err error
				m.amount, err = strconv.Atoi(m.Input.Value())
				if err != nil {
					m.msg = lipgloss.NewStyle().Foreground(Quaternary).Render("Not a valid amount.")
					return m, nil
				}

				// Check if user input is greater than 1
				if m.amount < 1 {
					m.msg = lipgloss.NewStyle().Foreground(Quaternary).Render("Requires an integer greater than 1.")
					return m, nil
				}

				m.msg = ""
				Models[Amount] = m
				return Models[Crafter].Update(nil)
			}
		}
	}

	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m Input) View() string {
	inputView := fmt.Sprintf(
		"Enter the Amount to Craft:\n\n%s",
		m.Input.View(),
	)

	msgView := "\n" + lipgloss.NewStyle().Bold(true).Foreground(Quaternary).Render(m.msg)

	return mainStyle.Render(lipgloss.JoinVertical(lipgloss.Left, titleView, inputView, msgView)) + "\n"
}
