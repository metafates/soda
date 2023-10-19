package soda

import (
	"context"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zyedidia/generic/stack"
	"golang.org/x/term"
)

type Option func(model *Model)

func WithKeyMap(keyMap KeyMap) Option {
	return func(model *Model) {
		model.keyMap = keyMap
	}
}

func WithOnError(onError OnErrorFunc) Option {
	return func(model *Model) {
		model.onError = onError
	}
}

// WithSize sets a fixed size for the Model.
func WithSize(size Size) Option {
	return func(model *Model) {
		model.size = size
	}
}

// WithAutoSize gets terminal dimensions and sets Model's size.
// It will also make the Model listen for incoming resize messages from temrinal.
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
		model.autoSize = true
	}
}

func WithStyleMap(styleMap StyleMap) Option {
	return func(model *Model) {
		model.styleMap = styleMap
	}
}

// WithSpinner sets a custom spinner.Spinner for Model
func WithSpinner(s spinner.Spinner) Option {
	return func(model *Model) {
		model.spinner.Spinner = s
	}
}

func WithSaveInitialStateToHistory(save bool) Option {
	return func(model *Model) {
		model.stateWrapper.SaveToHistory = save
	}
}

// New creates a new soda model with initial state
func New(state State, options ...Option) *Model {
	model := &Model{
		stateWrapper: stateWrapper{
			State:         state,
			SaveToHistory: true,
		},
		history: stack.New[stateWrapper](),
		onError: func(err error) tea.Cmd {
			return Notify(err.Error(), time.Second*3)
		},
		keyMap:   DefaultKeyMap(),
		styleMap: DefaultStyleMap(),
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

	model.keyMap.model = model
	model.spinner.Style = model.styleMap.Spinner

	return model
}
