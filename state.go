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

type State interface {
	Destroy()
	Backable() bool
	Intermediate() bool
	Resize(size Size) tea.Cmd

	Title() title.Title
	Subtitle() string

	KeyMap() help.KeyMap

	Init(mh ModelHandler) tea.Cmd
	Update(mh ModelHandler, msg tea.Msg) tea.Cmd
	View(mh ModelHandler) string
}
