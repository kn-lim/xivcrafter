package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

type UpdateSettings struct {
	// XIVCrafter settings

	DelayModel      textinput.Model // TODO
	delay           int             // TODO
	KeyDelayModel   textinput.Model // TODO
	keyDelay        int             // TODO
	StartPauseModel textinput.Model
	startPause      string
	StopModel       textinput.Model
	stop            string
	ConfirmModel    textinput.Model
	confirm         string
	CancelModel     textinput.Model
	cancel          string

	// Help model
	Help help.Model

	// Status message
	Msg string
}

func (m UpdateSettings) Init() tea.Cmd {
	return nil
}

func (m UpdateSettings) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit XIVCrafter
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Go back to previous textinput model or back to list model
		case "shift+tab", "ctrl+b":
			if m.StopModel.Focused() {
				m.StopModel.Blur()
				m.StartPauseModel.Focus()
				return m, textinput.Blink
			} else if m.ConfirmModel.Focused() {
				m.ConfirmModel.Blur()
				m.StopModel.Focus()
				return m, textinput.Blink
			} else if m.CancelModel.Focused() {
				m.CancelModel.Blur()
				m.ConfirmModel.Focus()
				return m, textinput.Blink
			}

			// At StartPause textinput model, therefore, exit
			return m, tea.Quit

		case "tab", "enter":
			if m.StartPauseModel.Focused() {
				// Use placeholder if blank
				value := m.StartPauseModel.Value()
				if value == "" {
					value = m.StartPauseModel.Placeholder
				}

				// Check if given name is valid
				if value == "" {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Start/Pause must not be blank.")
					return m, nil
				}
				if !utils.CheckValidKey(value) {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Start/Pause is not a valid key.")
					return m, nil
				}

				// Success
				m.Msg = ""
				m.startPause = value
				m.StartPauseModel.Blur()
				m.StopModel.Focus()
				return m, textinput.Blink
			} else if m.StopModel.Focused() {
				// Use placeholder if blank
				value := m.StopModel.Value()
				if value == "" {
					value = m.StopModel.Placeholder
				}

				// Check if given name is valid
				if value == "" {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Stop must not be blank.")
					return m, nil
				}
				if !utils.CheckValidKey(value) {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Stop is not a valid key.")
					return m, nil
				}

				// Success
				m.Msg = ""
				m.stop = value
				m.StopModel.Blur()
				m.ConfirmModel.Focus()
				return m, textinput.Blink
			} else if m.ConfirmModel.Focused() {
				// Use placeholder if blank
				value := m.ConfirmModel.Value()
				if value == "" {
					value = m.ConfirmModel.Placeholder
				}

				// Check if given name is valid
				if value == "" {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Confirm must not be blank.")
					return m, nil
				}
				if !utils.CheckValidKey(value) {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Confirm is not a valid key.")
					return m, nil
				}

				// Success
				m.Msg = ""
				m.confirm = value
				m.ConfirmModel.Blur()
				m.CancelModel.Focus()
				return m, textinput.Blink
			} else if m.CancelModel.Focused() {
				// Use placeholder if blank
				value := m.CancelModel.Value()
				if value == "" {
					value = m.CancelModel.Placeholder
				}

				// Check if given name is valid
				if value == "" {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Cancel must not be blank.")
					return m, nil
				}
				if !utils.CheckValidKey(value) {
					m.Msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Cancel is not a valid key.")
					return m, nil
				}

				// Success
				m.Msg = ""
				m.cancel = value
				m.CancelModel.Blur()

				utils.Log("Infow", "updating xivcrafter settings",
					"delay", m.delay,
					"key_delay", m.keyDelay,
					"start_pause", m.startPause,
					"stop", m.stop,
					"confirm", m.confirm,
					"cancel", m.cancel,
				)

				// Go back to list with new XIVCrafter settings
				return Models[Recipes].Update(*NewSettings(m.startPause, m.stop, m.confirm, m.cancel))
			}
		}

	case tea.WindowSizeMsg:
		WindowWidth = msg.Width
		WindowHeight = msg.Height
		h, v := utils.ListStyle.GetFrameSize()
		model := Models[Recipes].(*List)
		model.Recipes.SetSize(msg.Width-h, msg.Height-v)
		Models[Recipes] = model

	// Edit recipe
	case Settings:
		utils.Log("Infow", "updating xivcrafter settings")

		m.AddPlaceholders(msg)

		return m, nil
	}

	var cmd tea.Cmd
	if m.StartPauseModel.Focused() {
		m.StartPauseModel, cmd = m.StartPauseModel.Update(msg)
	} else if m.StopModel.Focused() {
		m.StopModel, cmd = m.StopModel.Update(msg)
	} else if m.ConfirmModel.Focused() {
		m.ConfirmModel, cmd = m.ConfirmModel.Update(msg)
	} else if m.CancelModel.Focused() {
		m.CancelModel, cmd = m.CancelModel.Update(msg)
	}
	return m, cmd
}

func (m UpdateSettings) View() string {
	_, v := utils.MainStyle.GetFrameSize()

	mainView := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Left, utils.TitleView, utils.StatusStyle.Render(m.Msg)),
		"",
		lipgloss.NewStyle().Render("Enter XIVCrafter Settings:\n"),
		m.StartPauseModel.View(),
		m.StopModel.View(),
		m.ConfirmModel.View(),
		m.CancelModel.View(),
	)
	mainView = lipgloss.NewStyle().Height(WindowHeight - v - 1).Render(mainView)

	return utils.MainStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		mainView,
		m.Help.View(updateKeys),
	))
}

// NewUpdateSettings returns a pointer to an UpdateSettings struct
func NewUpdateSettings() *UpdateSettings {
	model := &UpdateSettings{
		StartPauseModel: textinput.New(),
		StopModel:       textinput.New(),
		ConfirmModel:    textinput.New(),
		CancelModel:     textinput.New(),
		Help:            help.New(),
	}

	// Defaults
	model.StartPauseModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Foreground(utils.Secondary).Render("Start/Pause"))
	model.StopModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Foreground(utils.Secondary).Render("Stop"))
	model.ConfirmModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Foreground(utils.Secondary).Render("Confirm"))
	model.CancelModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Foreground(utils.Secondary).Render("Cancel"))

	model.StartPauseModel.Focus()
	return model
}

// AddPlaceholders updates the textinput.Model Placeholder fields to show the value from Item
func (m *UpdateSettings) AddPlaceholders(settings Settings) {
	m.StartPauseModel.Placeholder = settings.startPause
	m.StopModel.Placeholder = settings.stop
	m.ConfirmModel.Placeholder = settings.confirm
	m.CancelModel.Placeholder = settings.cancel
}

type Settings struct {
	startPause string
	stop       string
	confirm    string
	cancel     string
}

// NewSettings returns a pointer to a NewSettings struct
func NewSettings(startPause string, stop string, confirm string, cancel string) *Settings {
	return &Settings{
		startPause,
		stop,
		confirm,
		cancel,
	}
}
