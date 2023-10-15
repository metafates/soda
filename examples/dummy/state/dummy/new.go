package dummy

import "github.com/charmbracelet/bubbles/key"

func New() *State {
	return &State{
		keyMap: keyMap{
			SendNotification: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "send notification")),
			ToggleSpinner:    key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "toggle spinner")),
		},
	}
}
