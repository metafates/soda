package soda

import "time"

type backMsg struct {
	Steps int
}

type backToRootMsg struct{}

type pushStateMsg struct {
	State State
}

type replaceStateMsg struct {
	State State
}

type notificationMsg struct {
	Message  string
	Duration time.Duration
}

type startSpinnerMsg struct{}
type stopSpinnerMsg struct{}

type notificationTimeoutMsg struct{}
