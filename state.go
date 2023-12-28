package soda

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda/title"
)

type stateWrapper struct {
	State

	SaveToHistory bool
}

type State interface {
	Destroy()
	Focused() bool
	SetSize(size Size) tea.Cmd

	Title() title.Title
	Subtitle() string
	Status() string

	KeyMap() help.KeyMap

	Init(ctx context.Context) tea.Cmd
	Update(ctx context.Context, msg tea.Msg) tea.Cmd
	View() string
}
