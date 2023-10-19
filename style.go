package soda

import "github.com/charmbracelet/lipgloss"

type StyleMap struct {
	Global       lipgloss.Style
	Header       lipgloss.Style
	Subtitle     lipgloss.Style
	Status       lipgloss.Style
	Spinner      lipgloss.Style
	Notification lipgloss.Style
	HelpBar      lipgloss.Style
}

func DefaultStyleMap() StyleMap {
	return StyleMap{
		Global: lipgloss.
			NewStyle().
			Padding(0, 1),
		Header: lipgloss.
			NewStyle().
			Padding(0, 0, 1, 0),
		Subtitle: lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
		Status: lipgloss.
			NewStyle().
			Padding(0, 1),
		Spinner: lipgloss.
			NewStyle().
			Padding(0, 1),
		Notification: lipgloss.
			NewStyle().
			Italic(true).
			Padding(0, 1).
			AlignHorizontal(lipgloss.Right),
		HelpBar: lipgloss.
			NewStyle().
			Padding(1, 1, 0, 0),
	}
}
