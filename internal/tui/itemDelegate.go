package tui

import (
	"github.com/charmbracelet/bubbles/list"
)

func NewItemDelegate() list.ItemDelegate {
	// Create a new default delegate
	d := list.NewDefaultDelegate()

	// Change colors
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(Secondary).BorderLeftForeground(Secondary)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy()

	d.Styles.NormalTitle = d.Styles.NormalTitle.Foreground(Tertiary).BorderLeftForeground(Tertiary)
	d.Styles.NormalDesc = d.Styles.NormalTitle.Copy()

	return d
}
