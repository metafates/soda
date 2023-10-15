package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func Back() tea.Msg {
	return BackMsg{Steps: 1}
}

func BackN(steps int) tea.Cmd {
	return func() tea.Msg {
		return BackMsg{Steps: steps}
	}
}

func Error(err error) tea.Cmd {
	return func() tea.Msg {
		return err
	}
}

func PushState(state State) tea.Cmd {
	return func() tea.Msg {
		return PushStateMsg{State: state}
	}
}

func ReplaceState(state State) tea.Cmd {
	return func() tea.Msg {
		return ReplaceStateMsg{State: state}
	}
}

func Notify(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return NotificationMsg{Message: message, Duration: duration}
	}
}
