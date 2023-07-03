package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func NewItemDelegate() list.ItemDelegate {
	// Create a new default delegate
	d := list.NewDefaultDelegate()

	// Change colors
	c := lipgloss.Color("#6f03fc")
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(c).BorderLeftForeground(c)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy() // reuse the title style here

	return d
}
