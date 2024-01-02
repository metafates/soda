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
	keys = append(keys, k.stateKeyMap().ShortHelp()...)

	return keys
}

func (k KeyMap) FullHelp() [][]key.Binding {
	sections := [][]key.Binding{k.ShortHelp()}

	sections = append(sections, k.stateKeyMap().FullHelp()...)

	return sections
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
		Back: key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
		Help: key.NewBinding(key.WithKeys("ctrl+?"), key.WithHelp("ctrl+?", "help")),
	}
}
