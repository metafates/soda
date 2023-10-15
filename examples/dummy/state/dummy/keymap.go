package dummy

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*keyMap)(nil)

type keyMap struct {
	SendNotification key.Binding
	ToggleSpinner    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.SendNotification,
		k.ToggleSpinner,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
