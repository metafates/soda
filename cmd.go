package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

const DefaultNotificationTimeout = time.Second * 3

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

func SendCmd(supplier func() (tea.Cmd, error)) tea.Cmd {
	return func() tea.Msg {
		cmd, err := supplier()
		if err != nil {
			return err
		}

		if cmd == nil {
			return nil
		}

		return cmd()
	}
}

func ReplaceState(state State) tea.Cmd {
	return func() tea.Msg {
		return replaceStateMsg{State: state}
	}
}

// Notify sends a notification to the Model with default timeout
func Notify(message string) tea.Cmd {
	return NotifyTimeout(message, DefaultNotificationTimeout)
}

// NotifyTimeout sends a notification to the Model with custom timeout
func NotifyTimeout(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return notificationMsg{Message: message, Duration: duration}
	}
}

func Do(description string, cmd tea.Cmd) tea.Cmd {
	ID := uuid.NewString()

	return tea.Sequence(
		Notify(description),
		startSpinner(ID),
		cmd,
		stopSpinner(ID),
	)
}

// StartSpinner starts the Model's spinner
func startSpinner(ID string) tea.Cmd {
	return func() tea.Msg {
		return spinnerMsg{
			ID:   ID,
			stop: false,
		}
	}
}

// StopSpinner stops the Model's spinner
func stopSpinner(ID string) tea.Cmd {
	return func() tea.Msg {
		return spinnerMsg{
			ID:   ID,
			stop: true,
		}
	}
}
