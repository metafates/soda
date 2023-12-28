package soda

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zyedidia/generic/stack"
)

type Model struct {
	state   stateWrapper
	history *stack.Stack[stateWrapper]

	size Size

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (m *Model) stateSize() Size {
	// TODO
	return m.size
}

func (m *Model) cancel() {
	m.ctxCancel()
	m.ctx, m.ctxCancel = context.WithCancel(context.Background())
}

func (m *Model) resize(size Size) tea.Cmd {
	m.size = size
	return m.setStateSize()
}

func (m *Model) setStateSize() tea.Cmd {
	return m.state.SetSize(m.stateSize())
}

func (m *Model) back(steps int) tea.Cmd {
	// do not pop the last state
	if m.history.Size() == 0 || steps <= 0 {
		return nil
	}

	m.cancel()
	for i := 0; i < steps && m.history.Size() > 0; i++ {
		m.state.Destroy()
		m.state = m.history.Pop()
	}

	return m.initState()
}

func (m *Model) initState() tea.Cmd {
	return tea.Sequence(
		m.state.Init(m.ctx),
		m.setStateSize(),
	)
}

func (m *Model) pushState(state stateWrapper) tea.Cmd {
	if m.state.SaveToHistory {
		m.history.Push(m.state)
	}

	m.state = state
	return m.initState()
}

func (m *Model) replaceState(state State) tea.Cmd {
	m.state.Destroy()
	m.state.State = state

	return m.initState()
}
