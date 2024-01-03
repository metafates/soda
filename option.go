package soda

import "github.com/charmbracelet/bubbles/spinner"

type Option func(*Model)

func WithKeyMap(keyMap KeyMap) Option {
	return func(model *Model) {
		model.keyMap = keyMap
	}
}

func WithOnError(onError OnError) Option {
	return func(model *Model) {
		model.onError = onError
	}
}

func WithSpinner(spinner spinner.Spinner) Option {
	return func(model *Model) {
		model.spinner.Spinner = spinner
	}
}

func WithStyles(styles Styles) Option {
	return func(model *Model) {
		model.styles = styles
	}
}

func WithMinSize(size Size) Option {
	return func(model *Model) {
		model.minSize = size
	}
}
