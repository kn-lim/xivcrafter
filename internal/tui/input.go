package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

type Input struct {
	// Get user input for the amount to craft
	Input textinput.Model

	// Help model
	Help help.Model

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
				// Use placeholder if blank
				value := m.Input.Value()
				if value == "" {
					value = m.Input.Placeholder
				}

				// Check if user input is a valid integer
				amount, err := strconv.Atoi(value)
				if err != nil {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Not a valid amount")
					return m, nil
				}

				// Check if user input is greater than 1
				if amount < 1 {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Requires an integer greater than 1")
					return m, nil
				}

				if utils.Logger != nil {
					utils.Logger.Printf("Amount to craft: %v\n", amount)
				}

				// Go to Progress model
				m.msg = ""
				return Models[Crafter].Update(amount)
			}
		}
	}

	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m Input) View() string {
	return utils.MainStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Left, utils.TitleView, utils.StatusStyle.Render(m.msg)),
		"",
		m.Input.View(),
		"\n\n\n\n",
		m.Help.View(inputKeys),
	))
}

// NewInput returns a pointer to an Input struct
func NewInput() *Input {
	model := &Input{
		Input: textinput.New(),
		Help:  help.New(),
	}

	// Defaults
	model.Input.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Foreground(utils.Secondary).Render("Enter the Amount to Craft"))
	model.Input.Placeholder = "1"

	model.Input.Focus()
	return model
}
