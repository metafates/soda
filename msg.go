package soda

import "time"

type BackMsg struct {
	Steps int
}

type BackToRootMsg struct{}

type PushStateMsg struct {
	State State
}

type ReplaceStateMsg struct {
	State State
}

type NotificationMsg struct {
	Message  string
	Duration time.Duration
}

type notificationTimeoutMsg struct{}
