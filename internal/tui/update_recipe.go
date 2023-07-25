package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

type UpdateRecipe struct {
	// Recipe settings
	NameModel textinput.Model
	name      string

	// Consumables
	FoodModel         textinput.Model
	food              string
	FoodDurationModel textinput.Model
	foodDuration      int
	PotionModel       textinput.Model
	potion            string

	// In-game hotkeys
	Macro1Model         textinput.Model
	macro1              string
	Macro1DurationModel textinput.Model
	macro1Duration      int
	Macro2Model         textinput.Model
	macro2              string
	Macro2DurationModel textinput.Model
	macro2Duration      int
	Macro3Model         textinput.Model
	macro3              string
	Macro3DurationModel textinput.Model
	macro3Duration      int

	// Helpers
	msg string
}

func (m UpdateRecipe) Init() tea.Cmd {
	return nil
}

func (m UpdateRecipe) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit XIVCrafter
		case "ctrl+c", "q":
			return m, tea.Quit

		// Go back to recipes model
		case "esc":
			return Models[Recipes].Update(nil)

		// Go back to previous textinput model or back to list model
		case "shift+tab", "ctrl+b":
			if m.FoodModel.Focused() {
				m.FoodModel.Blur()
				m.NameModel.Focus()
				return m, textinput.Blink
			} else if m.FoodDurationModel.Focused() {
				m.FoodDurationModel.Blur()
				m.FoodModel.Focus()
				return m, textinput.Blink
			} else if m.PotionModel.Focused() {
				m.PotionModel.Blur()
				m.FoodDurationModel.Focus()
				return m, textinput.Blink
			} else if m.Macro1Model.Focused() {
				m.Macro1Model.Blur()
				m.PotionModel.Focus()
				return m, textinput.Blink
			} else if m.Macro1DurationModel.Focused() {
				m.Macro1DurationModel.Blur()
				m.Macro1Model.Focus()
				return m, textinput.Blink
			} else if m.Macro2Model.Focused() {
				m.Macro2Model.Blur()
				m.Macro1DurationModel.Focus()
				return m, textinput.Blink
			} else if m.Macro2DurationModel.Focused() {
				m.Macro2DurationModel.Blur()
				m.Macro2Model.Focus()
				return m, textinput.Blink
			} else if m.Macro3Model.Focused() {
				m.Macro3Model.Blur()
				m.Macro2DurationModel.Focus()
				return m, textinput.Blink
			} else if m.Macro3DurationModel.Focused() {
				m.Macro3DurationModel.Blur()
				m.Macro3Model.Focus()
				return m, textinput.Blink
			}

			// At Name textinput model, therefore, go back to List model
			return Models[Recipes].Update(nil)

		case "tab", "enter":
			if m.NameModel.Focused() {
				// Use placeholder if blank
				value := m.NameModel.Value()
				if value == "" {
					value = m.NameModel.Placeholder
				}

				// Check if given name is valid
				if value == "" {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Name must not be blank.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.name = value
				m.NameModel.Blur()
				m.FoodModel.Focus()
				return m, textinput.Blink
			} else if m.FoodModel.Focused() {
				// Use placeholder if blank
				value := m.FoodModel.Value()
				if value == "" {
					value = m.FoodModel.Placeholder
				}

				// Check if given food key is valid
				if !utils.CheckValidKey(value) {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Food is not a valid key.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.food = value
				m.FoodModel.Blur()
				m.FoodDurationModel.Focus()
				return m, textinput.Blink
			} else if m.FoodDurationModel.Focused() {
				// Use placeholder if blank
				value := m.FoodDurationModel.Value()
				if value == "" {
					value = m.FoodDurationModel.Placeholder
				}

				// Check if given food duration is valid
				switch value {
				case "30", "40", "45":
					// Success
					m.msg = ""
					m.foodDuration, _ = strconv.Atoi(value)
					m.FoodDurationModel.Blur()
					m.PotionModel.Focus()
					return m, textinput.Blink
				default:
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Food Duration must be either 30, 40 or 45")
					return m, nil
				}
			} else if m.PotionModel.Focused() {
				// Use placeholder if blank
				value := m.PotionModel.Value()
				if value == "" {
					value = m.PotionModel.Placeholder
				}

				// Check if given food key is valid
				if !utils.CheckValidKey(value) {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Potion is not a valid key.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.potion = value
				m.PotionModel.Blur()
				m.Macro1Model.Focus()
				return m, textinput.Blink
			} else if m.Macro1Model.Focused() {
				// Use placeholder if blank
				value := m.Macro1Model.Value()
				if value == "" {
					value = m.Macro1Model.Placeholder
				}

				// Check if given macro 1 key is valid
				if value == "" {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 1 must not be blank.")
					return m, nil
				}
				if !utils.CheckValidKey(value) {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 1 is not a valid key.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.macro1 = value
				m.Macro1Model.Blur()
				m.Macro1DurationModel.Focus()
				return m, textinput.Blink
			} else if m.Macro1DurationModel.Focused() {
				// Use placeholder if blank
				value := m.Macro1DurationModel.Value()
				if value == "" {
					value = m.Macro1DurationModel.Placeholder
				}

				// Check if given macro 1 duration is valid
				macro1Duration, err := strconv.Atoi(value)
				if err != nil || macro1Duration < 1 {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 1 Duration is not a valid duration.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.macro1Duration = macro1Duration
				m.Macro1DurationModel.Blur()
				m.Macro2Model.Focus()
				return m, textinput.Blink
			} else if m.Macro2Model.Focused() {
				// Use placeholder if blank
				value := m.Macro2Model.Value()
				if value == "" {
					value = m.Macro2Model.Placeholder
				}

				// Check if given macro 2 key is valid
				if !utils.CheckValidKey(value) {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 2 is not a valid key.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.macro2 = value
				m.Macro2Model.Blur()
				m.Macro2DurationModel.Focus()
				return m, textinput.Blink
			} else if m.Macro2DurationModel.Focused() {
				// Use placeholder if blank
				value := m.Macro2DurationModel.Value()
				if value == "" {
					value = m.Macro2DurationModel.Placeholder
				}

				// Check if given macro 2 duration is valid
				macro2Duration, err := strconv.Atoi(value)
				if err != nil || macro2Duration < 1 {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 2 Duration is not a valid duration.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.macro2Duration = macro2Duration
				m.Macro2DurationModel.Blur()
				m.Macro3Model.Focus()
				return m, textinput.Blink
			} else if m.Macro3Model.Focused() {
				// Use placeholder if blank
				value := m.Macro3Model.Value()
				if value == "" {
					value = m.Macro3Model.Placeholder
				}

				// Check if given macro 3 key is valid
				if !utils.CheckValidKey(value) {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 3 is not a valid key.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.macro3 = value
				m.Macro3Model.Blur()
				m.Macro3DurationModel.Focus()
				return m, textinput.Blink
			} else if m.Macro3DurationModel.Focused() {
				// Use placeholder if blank
				value := m.Macro3DurationModel.Value()
				if value == "" {
					value = m.Macro3DurationModel.Placeholder
				}

				// Check if given macro 3 duration is valid
				macro3Duration, err := strconv.Atoi(value)
				if err != nil || macro3Duration < 1 {
					m.msg = lipgloss.NewStyle().Foreground(utils.Red).Render("Macro 3 Duration is not a valid duration.")
					return m, nil
				}

				// Success
				m.msg = ""
				m.macro3Duration = macro3Duration
				m.Macro3DurationModel.Blur()

				recipe := Item{
					Name:           m.name,
					Food:           m.food,
					FoodDuration:   m.foodDuration,
					Potion:         m.potion,
					Macro1:         m.macro1,
					Macro1Duration: m.macro1Duration,
					Macro2:         m.macro2,
					Macro2Duration: m.macro2Duration,
					Macro3:         m.macro3,
					Macro3Duration: m.macro3Duration,
				}

				if utils.Logger != nil {
					utils.Logger.Printf("Creating Recipe: %v\n", recipe)
				}

				// Go back to list with new recipe
				return Models[Recipes].Update(recipe)
			}
		}

	// Edit recipe
	case Item:
		if utils.Logger != nil {
			utils.Logger.Printf("Editing recipe %s\n", msg.Name)
		}

		m.AddPlaceholders(msg)

		return m, nil

	// From UpdateSettings model
	case Settings:
		//
	}

	var cmd tea.Cmd
	if m.NameModel.Focused() {
		m.NameModel, cmd = m.NameModel.Update(msg)
	} else if m.FoodModel.Focused() {
		m.FoodModel, cmd = m.FoodModel.Update(msg)
	} else if m.FoodDurationModel.Focused() {
		m.FoodDurationModel, cmd = m.FoodDurationModel.Update(msg)
	} else if m.PotionModel.Focused() {
		m.PotionModel, cmd = m.PotionModel.Update(msg)
	} else if m.Macro1Model.Focused() {
		m.Macro1Model, cmd = m.Macro1Model.Update(msg)
	} else if m.Macro1DurationModel.Focused() {
		m.Macro1DurationModel, cmd = m.Macro1DurationModel.Update(msg)
	} else if m.Macro2Model.Focused() {
		m.Macro2Model, cmd = m.Macro2Model.Update(msg)
	} else if m.Macro2DurationModel.Focused() {
		m.Macro2DurationModel, cmd = m.Macro2DurationModel.Update(msg)
	} else if m.Macro3Model.Focused() {
		m.Macro3Model, cmd = m.Macro3Model.Update(msg)
	} else if m.Macro3DurationModel.Focused() {
		m.Macro3DurationModel, cmd = m.Macro3DurationModel.Update(msg)
	}
	return m, cmd
}

func (m UpdateRecipe) View() string {
	return mainStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Left, titleView, lipgloss.NewStyle().Padding(1, 0, 0, 4).Bold(true).Render(m.msg)),
		"",
		lipgloss.NewStyle().Bold(true).Render("Enter Recipe Settings:\n"),
		m.NameModel.View(),
		m.FoodModel.View(),
		m.FoodDurationModel.View(),
		m.PotionModel.View(),
		m.Macro1Model.View(),
		m.Macro1DurationModel.View(),
		m.Macro2Model.View(),
		m.Macro2DurationModel.View(),
		m.Macro3Model.View(),
		m.Macro3DurationModel.View(),
	))
}

func NewUpdateRecipe() *UpdateRecipe {
	model := &UpdateRecipe{
		NameModel:           textinput.New(),
		FoodModel:           textinput.New(),
		FoodDurationModel:   textinput.New(),
		PotionModel:         textinput.New(),
		Macro1Model:         textinput.New(),
		Macro1DurationModel: textinput.New(),
		Macro2Model:         textinput.New(),
		Macro2DurationModel: textinput.New(),
		Macro3Model:         textinput.New(),
		Macro3DurationModel: textinput.New(),
	}

	// Defaults
	model.NameModel.CharLimit = MaxNameLength
	model.NameModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Name"))
	model.FoodModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Food"))
	model.FoodDurationModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Food Duration"))
	model.FoodDurationModel.Placeholder = "30"
	model.PotionModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Potion"))
	model.Macro1Model.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Macro 1"))
	model.Macro1DurationModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Macro 1 Duration"))
	model.Macro1DurationModel.Placeholder = "1"
	model.Macro2Model.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Macro 2"))
	model.Macro2DurationModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Macro 2 Duration"))
	model.Macro2DurationModel.Placeholder = "1"
	model.Macro3Model.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Macro 3"))
	model.Macro3DurationModel.Prompt = fmt.Sprintf("%s: ", lipgloss.NewStyle().Bold(true).Render("Macro 3 Duration"))
	model.Macro3DurationModel.Placeholder = "1"

	model.NameModel.Focus()
	return model
}

func (m *UpdateRecipe) AddPlaceholders(item Item) {
	m.NameModel.Placeholder = item.Name
	m.FoodModel.Placeholder = item.Food
	m.FoodDurationModel.Placeholder = strconv.Itoa(item.FoodDuration)
	m.PotionModel.Placeholder = item.Potion
	m.Macro1Model.Placeholder = item.Macro1
	m.Macro1DurationModel.Placeholder = strconv.Itoa(item.Macro1Duration)
	m.Macro2Model.Placeholder = item.Macro2
	m.Macro2DurationModel.Placeholder = strconv.Itoa(item.Macro2Duration)
	m.Macro3Model.Placeholder = item.Macro3
	m.Macro3DurationModel.Placeholder = strconv.Itoa(item.Macro3Duration)
}
