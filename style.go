package soda

import "github.com/charmbracelet/lipgloss"

type StyleMap struct {
	TitleBar lipgloss.Style
	Subtitle lipgloss.Style
	HelpBar  lipgloss.Style
}

func DefaultStyleMap() StyleMap {
	return StyleMap{
		TitleBar: lipgloss.
			NewStyle().
			Padding(0, 0, 1, 2),
		Subtitle: lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
		HelpBar: lipgloss.
			NewStyle().
			Padding(0, 1),
	}
}
