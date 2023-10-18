package soda

import "github.com/charmbracelet/lipgloss"

type StyleMap struct {
	TitleBar     lipgloss.Style
	Subtitle     lipgloss.Style
	Notification lipgloss.Style
	HelpBar      lipgloss.Style
}

func DefaultStyleMap() StyleMap {
	return StyleMap{
		TitleBar: lipgloss.
			NewStyle().
			//Padding(0, 0, 1, 2),
			Padding(0, 0, 1, 0),
		Subtitle: lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
		Notification: lipgloss.NewStyle().Italic(true),
		HelpBar: lipgloss.
			NewStyle().
			Padding(1, 1, 0, 0),
	}
}
