package tui

import "github.com/charmbracelet/bubbles/key"

type listKeyMap struct {
	enter          key.Binding
	changeSettings key.Binding
	newRecipe      key.Binding
	editRecipe     key.Binding
}

var listKeys = listKeyMap{
	enter:          key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	changeSettings: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "change settings")),
	newRecipe:      key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new recipe")),
	editRecipe:     key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit recipe")),
}

type inputKeyMap struct {
	// Key binding to submit the user input
	enter key.Binding

	// XIVCrafter standard key bindings
	back key.Binding
	quit key.Binding
}

func (k inputKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.enter, k.back, k.quit}
}

func (k inputKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.enter}, {k.back, k.quit}}
}

var inputKeys = inputKeyMap{
	enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	back:  key.NewBinding(key.WithKeys("esc", "b"), key.WithHelp("esc/b", "back")),
	quit:  key.NewBinding(key.WithKeys("ctrl+c", "q"), key.WithHelp("q", "quit")),
}

type progressKeyMap struct {
	// XIVCrafter standard key bindings
	back key.Binding
	quit key.Binding
}

func (k progressKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.back, k.quit}
}

func (k progressKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.back, k.quit}}
}

var progressKeys = progressKeyMap{
	back: key.NewBinding(key.WithKeys("esc", "b"), key.WithHelp("esc/b", "back")),
	quit: key.NewBinding(key.WithKeys("ctrl+c", "q"), key.WithHelp("q", "quit")),
}
