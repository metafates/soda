package soda

import (
	"context"
	"math"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda/title"
)

type stateWrapper struct {
	State         State
	SaveToHistory bool
}

type ModelHandler interface {
	StateSize() Size
	Context() context.Context
	SpinnerActive() bool
}

type Size struct {
	Width, Height int
}

func (s Size) SplitHorizotnal(size Size, leftRatio float64) (Size, Size) {
	width := float64(size.Width)
	left := width * leftRatio
	right := width - left

	return Size{
			Width:  int(math.Round(left)),
			Height: size.Height,
		}, Size{
			Width:  int(math.Round(right)),
			Height: size.Height,
		}
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

	// Subtitle is displayed under the title.
	// If empty string is given subtitle won't be rendered
	Subtitle() string

	// Status is displayed next to the title
	Status() string

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
