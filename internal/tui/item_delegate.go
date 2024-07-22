package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/kn-lim/xivcrafter/internal/utils"
)

func NewItemDelegate() list.ItemDelegate {
	// Create a new default delegate
	d := list.NewDefaultDelegate()

	// Change colors
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(utils.Secondary).BorderLeftForeground(utils.Secondary)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle

	return d
}
