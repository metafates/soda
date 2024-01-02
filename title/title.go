package title

import "github.com/charmbracelet/lipgloss"

type Option func(*Title)

func WithBackground(color lipgloss.Color) Option {
	return func(title *Title) {
		title.background = color
	}
}

func WithForeground(color lipgloss.Color) Option {
	return func(title *Title) {
		title.foreground = color
	}
}

func New(text string, options ...Option) Title {
	title := Title{
		text:       text,
		background: lipgloss.Color("#EB5E28"),
		foreground: lipgloss.Color("#252422"),
	}

	for _, option := range options {
		option(&title)
	}

	return title
}

type Title struct {
	text                   string
	background, foreground lipgloss.Color
}

func (t Title) String() string {
	return t.text
}

func (t Title) Text() string {
	return t.text
}

func (t Title) Style() lipgloss.Style {
	return lipgloss.
		NewStyle().
		Background(t.background).
		Foreground(t.foreground).
		Bold(true).
		Padding(0, 1)
}

func (t Title) Render(parents ...lipgloss.Style) string {
	style := t.Style()

	for _, parent := range parents {
		style.Inherit(parent)
	}

	return style.Render(t.text)
}
