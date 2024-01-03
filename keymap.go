package soda

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func NewKeyMap() KeyMap {
	return KeyMap{
		Back:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
		Quit:     key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
		ShowHelp: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	}
}

type KeyMap struct {
	Back, Quit, ShowHelp key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Back,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}

func (k KeyMap) with(other help.KeyMap) combinedKeyMap {
	return combinedKeyMap{
		overlays: []help.KeyMap{k, other},
	}
}

var _ help.KeyMap = (*combinedKeyMap)(nil)

type combinedKeyMap struct {
	overlays []help.KeyMap
}

func (c combinedKeyMap) ShortHelp() []key.Binding {
	var bindings []key.Binding

	for _, overlay := range c.overlays {
		bindings = append(bindings, overlay.ShortHelp()...)
	}

	return bindings
}

func (c combinedKeyMap) FullHelp() [][]key.Binding {
	var groups [][]key.Binding

	for _, overlay := range c.overlays {
		groups = append(groups, overlay.FullHelp()...)
	}

	return groups
}

func (c combinedKeyMap) With(other help.KeyMap) combinedKeyMap {
	return combinedKeyMap{
		overlays: append(c.overlays, other),
	}
}
