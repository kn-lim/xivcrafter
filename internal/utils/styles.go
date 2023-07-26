package utils

import "github.com/charmbracelet/lipgloss"

const (
	// Colors

	Primary   = lipgloss.Color("12")
	Secondary = lipgloss.Color("14")
	Tertiary  = lipgloss.Color("13")

	// Text Colors

	Gray   = lipgloss.Color("8")
	Red    = lipgloss.Color("9")
	Green  = lipgloss.Color("10")
	Yellow = lipgloss.Color("11")
)

var (
	// Styles

	TitleStyle      = lipgloss.NewStyle().Padding(0, 1).Background(Primary)
	TitleView       = TitleStyle.Render("XIVCrafter")
	ListStyle       = lipgloss.NewStyle().Margin(1, 1)
	MainStyle       = lipgloss.NewStyle().Margin(1, 3)
	StatusStyle     = lipgloss.NewStyle().PaddingLeft(6)
	ListStatusStyle = lipgloss.NewStyle().PaddingLeft(4)
)
