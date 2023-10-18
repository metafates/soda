package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Back pops one state back
func Back() tea.Msg {
	return backMsg{Steps: 1}
}

// BackN pops n states back
func BackN(steps int) tea.Cmd {
	return func() tea.Msg {
		return backMsg{Steps: steps}
	}
}

// Error sends error to the Model
func Error(err error) tea.Cmd {
	return func() tea.Msg {
		return err
	}
}

func PushState(state State, save bool) tea.Cmd {
	return func() tea.Msg {
		return pushStateMsg{State: state, Save: save}
	}
}

func PushStateFunc(stateSupplier func() (State, error), save bool) tea.Cmd {
	return func() tea.Msg {
		state, err := stateSupplier()
		if err != nil {
			return err
		}

		return pushStateMsg{
			State: state,
			Save:  save,
		}
	}
}

func ReplaceState(state State) tea.Cmd {
	return func() tea.Msg {
		return replaceStateMsg{State: state}
	}
}

// Notify sends a notification to the Model
func Notify(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return notificationMsg{Message: message, Duration: duration}
	}
}

// StartSpinner starts the Model's spinner
func StartSpinner() tea.Msg {
	return startSpinnerMsg{}
}

// StopSpinner stops the Model's spinner
func StopSpinner() tea.Msg {
	return stopSpinnerMsg{}
}
