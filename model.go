package soda

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"github.com/zyedidia/generic/stack"
)

type OnErrorFunc func(err error) tea.Cmd

var _ ModelHandler = (*Model)(nil)

type Model struct {
	showSpinner bool
	spinnersIDs map[string]struct{}
	autoSize    bool

	spinner      spinner.Model
	stateWrapper stateWrapper

	history  *stack.Stack[stateWrapper]
	onError  OnErrorFunc
	keyMap   KeyMap
	styleMap StyleMap
	size     Size
	help     help.Model

	notification      string
	notificationTimer *time.Timer

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (m *Model) state() State {
	return m.stateWrapper.State
}

func (m *Model) StateSize() Size {
	size := m.size

	size.Height -= 3

	if m.help.ShowAll {
		size.Height -= lipgloss.Height(m.help.View(m.keyMap))
	} else {
		size.Height--
	}

	if m.state().Subtitle() != "" {
		size.Height -= 2
	}

	size.Width -= m.styleMap.Global.GetHorizontalFrameSize()
	size.Height -= m.styleMap.Global.GetVerticalFrameSize()

	return size
}

func (m *Model) SpinnerActive() bool {
	return m.showSpinner
}

func (m *Model) Context() context.Context {
	return m.ctx
}

func (m *Model) Init() tea.Cmd {
	return m.initState()
}

func (m *Model) initState() tea.Cmd {
	return tea.Sequence(m.state().Init(m), m.resizeState())
}

func (m *Model) View() string {
	const newline = "\n"

	var titleBuilder strings.Builder

	titleBuilder.WriteString(m.state().Title().Render(lipgloss.NewStyle().MaxWidth(m.size.Width / 2)))

	if m.showSpinner {
		titleBuilder.WriteString(m.spinner.View())
	}

	if status := m.state().Status(); status != "" {
		titleBuilder.WriteString(m.styleMap.Status.Render(status))
	}

	if m.notification != "" {
		width := m.size.Width - lipgloss.Width(titleBuilder.String())
		titleBuilder.WriteString(m.styleMap.Notification.Width(width).Render(m.notification))
	}

	if subtitle := m.state().Subtitle(); subtitle != "" {
		subtitle = m.styleMap.Subtitle.Render(subtitle)

		titleBuilder.WriteString(newline + newline)
		titleBuilder.WriteString(m.styleMap.Subtitle.Render(subtitle))
	}

	header := m.styleMap.Header.Render(titleBuilder.String())

	state := wordwrap.String(m.state().View(m), m.size.Width)
	help := m.styleMap.HelpBar.Render(m.help.View(m.keyMap))

	headerHeight := lipgloss.Height(header)
	stateHeight := lipgloss.Height(state)
	helpHeight := lipgloss.Height(help)

	diff := m.size.Height - headerHeight - stateHeight - helpHeight - m.styleMap.Global.GetVerticalFrameSize()

	var filler string
	if diff > 0 {
		filler = strings.Repeat(newline, diff)
	}

	var sb strings.Builder

	sb.Grow(len(header))
	sb.Grow(len(state))
	sb.Grow(len(filler))
	sb.Grow(len(help))
	sb.Grow(len(newline) * 2)

	sb.WriteString(header)
	sb.WriteString(newline)
	sb.WriteString(state)
	sb.WriteString(filler)
	sb.WriteString(newline)
	sb.WriteString(help)

	return m.styleMap.Global.Render(sb.String())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.autoSize {
			return m, nil
		}

		cmd := m.resize(Size{
			Width:  msg.Width,
			Height: msg.Height,
		})

		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.Back) && m.state().Backable():
			return m, m.back(1)
		case key.Matches(msg, m.keyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
			cmd := m.resizeState()
			return m, cmd
		}
	case notificationMsg:
		cmd := m.setNotification(msg.Message, msg.Duration)
		return m, cmd
	case notificationTimeoutMsg:
		m.hideNotification()
		return m, nil
	case spinnerMsg:
		if msg.stop {
			_, ok := m.spinnersIDs[msg.ID]
			if ok {
				delete(m.spinnersIDs, msg.ID)

				if len(m.spinnersIDs) == 0 {
					return m, m.stopSpinner
				}
			}

			return m, nil
		}

		m.spinnersIDs[msg.ID] = struct{}{}
		return m, m.startSpinner
	case backMsg:
		return m, m.back(msg.Steps)
	case backToRootMsg:
		return m, m.back(m.history.Size())
	case pushStateMsg:
		return m, m.pushState(msg.State, msg.Save)
	case replaceStateMsg:
		return m, m.replaceState(msg.State)
	case spinnerTickMsg:
		if !m.showSpinner {
			return m, nil
		}

		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg.Msg)

		return m, func() tea.Msg {
			var msg tea.Msg
			if cmd != nil {
				msg = cmd()
			}

			return spinnerTickMsg{Msg: msg}
		}
	case error:
		if errors.Is(msg, context.Canceled) || strings.Contains(msg.Error(), context.Canceled.Error()) {
			return m, nil
		}

		return m, m.onError(msg)
	}

	cmd := m.state().Update(m, msg)
	return m, cmd
}

func (m *Model) cancel() {
	m.ctxCancel()
	m.ctx, m.ctxCancel = context.WithCancel(context.Background())
}

func (m *Model) resize(size Size) tea.Cmd {
	m.size = size
	m.help.Width = size.Width
	return m.resizeState()
}

func (m *Model) resizeState() tea.Cmd {
	return m.state().Resize(m.StateSize())
}

func (m *Model) back(steps int) tea.Cmd {
	// do not pop the last state
	if m.history.Size() == 0 || steps <= 0 {
		return NotifyTimeout("Can't go back", 400*time.Millisecond)
	}

	m.cancel()
	for i := 0; i < steps && m.history.Size() > 0; i++ {
		m.state().Destroy()
		m.stateWrapper = m.history.Pop()
	}

	return m.initState()
}

func (m *Model) pushState(state State, save bool) tea.Cmd {
	if m.stateWrapper.SaveToHistory {
		m.history.Push(m.stateWrapper)
	}

	m.stateWrapper = stateWrapper{
		State:         state,
		SaveToHistory: save,
	}

	return m.initState()
}

func (m *Model) replaceState(state State) tea.Cmd {
	m.state().Destroy()
	m.stateWrapper.State = state

	return m.initState()
}

func (m *Model) hideNotification() {
	m.notification = ""
	if m.notificationTimer != nil {
		m.notificationTimer.Stop()
	}
}

func (m *Model) setNotification(message string, duration time.Duration) tea.Cmd {
	m.notification = message

	if m.notificationTimer != nil {
		m.notificationTimer.Stop()
	}

	m.notificationTimer = time.NewTimer(duration)

	return func() tea.Msg {
		<-m.notificationTimer.C
		return notificationTimeoutMsg{}
	}
}

func (m *Model) startSpinner() tea.Msg {
	m.showSpinner = true

	return spinnerTickMsg{Msg: m.spinner.Tick()}
}

func (m *Model) stopSpinner() tea.Msg {
	m.showSpinner = false
	return nil
}
