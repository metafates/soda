package soda

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func Back() tea.Msg {
	return backMsg{Steps: 1}
}

func BackN(n int) tea.Cmd {
	if n < 0 {
		panic("n < 0")
	}

	return func() tea.Msg {
		return backMsg{Steps: n}
	}
}

func BackToRoot() tea.Msg {
	return backToRootMsg{}
}

func PushState(state State) tea.Cmd {
	return func() tea.Msg {
		return pushStateMsg{State: stateWrapper{
			State:         state,
			SaveToHistory: true,
		}}
	}
}

func PushTempState(state State) tea.Cmd {
	return func() tea.Msg {
		return pushStateMsg{State: stateWrapper{
			State:         state,
			SaveToHistory: false,
		}}
	}
}

func Notify(message string) tea.Cmd {
	return func() tea.Msg {
		return notificationMsg{Message: message}
	}
}

func NotifyWithDuration(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return notificationWithDurationMsg{
			notificationMsg: notificationMsg{
				Message: message,
			},
			Duration: duration,
		}
	}
}
