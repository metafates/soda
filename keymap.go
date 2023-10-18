package soda

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	model *Model

	Quit key.Binding
	Back key.Binding
	Help key.Binding
}

func (k KeyMap) stateKeyMap() help.KeyMap {
	return k.model.state().KeyMap()
}

func (k KeyMap) ShortHelp() []key.Binding {
	keys := []key.Binding{k.Quit}

	if k.model.history.Size() > 0 {
		keys = append(keys, k.Back)
	}

	keys = append(keys, k.Help)

	for _, key := range k.stateKeyMap().ShortHelp() {
		keys = append(keys, key)
	}

	return keys
}

func (k KeyMap) FullHelp() [][]key.Binding {
	sections := [][]key.Binding{k.ShortHelp()}

	for _, section := range k.stateKeyMap().FullHelp() {
		sections = append(sections, section)
	}

	return sections
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
		Back: key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
		Help: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	}
}
