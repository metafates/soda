package soda

import (
	"context"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zyedidia/generic/stack"
	"golang.org/x/term"
)

type Option func(model *Model)

func WithKeyMap(keyMap KeyMap) Option {
	return func(model *Model) {
		model.KeyMap = keyMap
	}
}

func WithOnError(onError OnErrorFunc) Option {
	return func(model *Model) {
		model.onError = onError
	}
}

func WithInitialSize(size Size) Option {
	return func(model *Model) {
		model.size = size
	}
}

func WithAutoSize() Option {
	return func(model *Model) {
		var (
			size Size
			err  error
		)

		size.Width, size.Height, err = term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			size.Width, size.Height = 80, 40
		}

		model.size = size
	}
}

func WithStyleMap(styleMap StyleMap) Option {
	return func(model *Model) {
		model.StyleMap = styleMap
	}
}

func WithSpinner(s spinner.Spinner) Option {
	return func(model *Model) {
		model.spinner.Spinner = s
	}
}

func New(state State, options ...Option) *Model {
	model := &Model{
		state:   state,
		history: stack.New[State](),
		onError: func(err error) tea.Cmd {
			return nil
		},
		KeyMap:   DefaultKeyMap(),
		StyleMap: DefaultStyleMap(),
		spinner:  spinner.New(),
		help:     help.New(),
		size: Size{
			Width:  80,
			Height: 40,
		},
	}
	model.ctx, model.ctxCancel = context.WithCancel(context.Background())

	for _, option := range options {
		option(model)
	}

	model.KeyMap.model = model

	return model
}
