package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func Back() tea.Msg {
	return backMsg{Steps: 1}
}

func BackN(steps int) tea.Cmd {
	return func() tea.Msg {
		return backMsg{Steps: steps}
	}
}

func Error(err error) tea.Cmd {
	return func() tea.Msg {
		return err
	}
}

func PushState(state State) tea.Cmd {
	return func() tea.Msg {
		return pushStateMsg{State: state}
	}
}

func ReplaceState(state State) tea.Cmd {
	return func() tea.Msg {
		return replaceStateMsg{State: state}
	}
}

func Notify(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return notificationMsg{Message: message, Duration: duration}
	}
}

func StartSpinner() tea.Msg {
	return startSpinnerMsg{}
}

func StopSpinner() tea.Msg {
	return stopSpinnerMsg{}
}
