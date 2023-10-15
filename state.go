package soda

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda/title"
)

type ModelHandler interface {
	StateSize() Size
	Context() context.Context
	SpinnerActive() bool
}

type Size struct {
	Width, Height int
}

// State that model operates
type State interface {
	// Destroy is called before the state is popped out from history
	Destroy()

	// Backable whether the state can be popped from history
	Backable() bool

	// Resize the state
	Resize(size Size) tea.Cmd

	// Title of the state
	Title() title.Title

	// Subtitle of the state. Displayed under the title.
	// If empty string is given subtitle won't be rendered
	Subtitle() string

	// KeyMap of the state
	KeyMap() help.KeyMap

	// Init is the first function that will be called. It returns an optional
	// initial command. To not perform an initial command return nil.
	Init(mh ModelHandler) tea.Cmd

	// Update is called when a message is received. Use it to inspect messages
	// and, in response, update the model and/or send a command.
	Update(mh ModelHandler, msg tea.Msg) tea.Cmd

	// View renders the state's UI, which is just a string. The view is
	// rendered after every Update.
	View(mh ModelHandler) string
}
