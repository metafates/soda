package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type backMsg struct {
	Steps int
}

type backToRootMsg struct{}

type pushStateMsg struct {
	State State
	Save  bool
}

type replaceStateMsg struct {
	State State
}

type notificationMsg struct {
	Message  string
	Duration time.Duration
}

type (
	startSpinnerMsg struct{}
	stopSpinnerMsg  struct{}
)

type notificationTimeoutMsg struct{}

type spinnerTickMsg struct {
	Msg tea.Msg
}
