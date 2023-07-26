package tui

import "github.com/charmbracelet/bubbles/key"

// Help model for List

// Additional key map for List model
type listKeyMap struct {
	// enter key
	enter key.Binding

	// change settings key
	change key.Binding

	// new recipe key
	new key.Binding

	// edit recipe key
	edit key.Binding

	// delete recipe key
	delete key.Binding
}

var listKeys = listKeyMap{
	enter:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	change: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "change settings")),
	new:    key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new")),
	edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit")),
	delete: key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "delete")),
}

// Help model for Update

// Key map for UpdateSettings and UpdateRecipe models
type updateKeyMap struct {
	// submit recipe key
	enter key.Binding

	// back to previous field key
	backField key.Binding

	// back to list model key
	backList key.Binding

	// quit key
	quit key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k updateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.enter, k.backField, k.quit}
}

// FullHelp returns keybindings for the expanded help view
func (k updateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.enter, k.backField}, {k.backList, k.quit}}
}

var updateKeys = updateKeyMap{
	enter:     key.NewBinding(key.WithKeys("tab", "enter"), key.WithHelp("tab/enter", "submit")),
	backField: key.NewBinding(key.WithKeys("shift+tab", "ctrl+b"), key.WithHelp("shift+tab/ctrl+b", "back")),
	backList:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back to recipe")),
	quit:      key.NewBinding(key.WithKeys("ctrl+c", "q"), key.WithHelp("ctrl+c/q", "quit")),
}

// Help model for Input

type inputKeyMap struct {
	// submit input key
	enter key.Binding

	// back to list model key
	back key.Binding

	// quit key
	quit key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k inputKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.enter, k.back, k.quit}
}

// FullHelp returns keybindings for the expanded help view
func (k inputKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.enter}, {k.back, k.quit}}
}

var inputKeys = inputKeyMap{
	enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	back:  key.NewBinding(key.WithKeys("esc", "b"), key.WithHelp("esc/b", "back")),
	quit:  key.NewBinding(key.WithKeys("ctrl+c", "q"), key.WithHelp("q", "quit")),
}

// Help model for Progress

type progressKeyMap struct {
	// back to input model key
	back key.Binding

	// quit key
	quit key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k progressKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.back, k.quit}
}

// FullHelp returns keybindings for the expanded help view
func (k progressKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.back, k.quit}}
}

var progressKeys = progressKeyMap{
	back: key.NewBinding(key.WithKeys("esc", "b"), key.WithHelp("esc/b", "back")),
	quit: key.NewBinding(key.WithKeys("ctrl+c", "q"), key.WithHelp("q", "quit")),
}
