package soda

import (
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zyedidia/generic/stack"
	"strings"
	"time"
)

var _ tea.Model = (*Model)(nil)

type notification struct {
	Message string
	Timer   *time.Timer
}

// OnError is the function that is called when any error occurs.
type OnError func(err error) tea.Cmd

type Model struct {
	minSize Size

	styles Styles

	state   stateWrapper
	history *stack.Stack[stateWrapper]

	onError OnError

	spinner     spinner.Model
	showSpinner bool

	size Size

	keyMap KeyMap

	help help.Model

	notificationDefaultDuration time.Duration
	notification                notification

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	if !m.isValidSize() {
		return m.viewInvalidSizeBanner()
	}

	stateSize := m.stateSize()

	stateView := m.state.View()
	stateView = lipgloss.Place(
		stateSize.Width,
		stateSize.Height,
		lipgloss.Left,
		lipgloss.Top,
		stateView,
	)

	view := lipgloss.JoinVertical(
		lipgloss.Left,
		m.viewHeader(),
		stateView,
		m.viewFooter(),
	)

	return view
}

func (m *Model) viewInvalidSizeBanner() string {
	banner := lipgloss.JoinVertical(
		lipgloss.Center,
		"Terminal size is too small:",
		fmt.Sprintf("Width = %d Height = %d", m.size.Width, m.size.Height),
		"",
		"Needed:",
		fmt.Sprintf("Width >= %d Height >= %d", m.minSize.Width, m.minSize.Height),
	)

	banner = lipgloss.NewStyle().Bold(true).Render(banner)

	return lipgloss.Place(
		m.size.Width,
		m.size.Height,
		lipgloss.Center,
		lipgloss.Center,
		banner,
	)
}

func (m *Model) viewHeader() string {
	var b strings.Builder

	b.Grow(200)

	title := m.state.Title().Render(lipgloss.NewStyle().MaxWidth(m.size.Width / 2))
	b.WriteString(title)

	if status := m.state.Status(); status != "" {
		b.WriteString(m.styles.Status.Render(status))
	}

	if m.notification.Message != "" {
		width := m.size.Width - lipgloss.Width(b.String())
		b.WriteString(m.styles.Notification.Copy().Width(width).Render(m.notification.Message))
	}

	if subtitle := m.state.Subtitle(); subtitle != "" {
		subtitle = m.styles.Subtitle.Render(subtitle)

		b.WriteString("\n\n")
		b.WriteString(m.styles.Subtitle.Render(subtitle))
	}

	header := m.styles.Header.Render(b.String())

	return header
}

func (m *Model) viewFooter() string {
	keyMap := m.keyMap.with(m.state.KeyMap())
	helpView := m.help.View(keyMap)

	footer := m.styles.Footer.Render(helpView)

	return footer
}

func (m *Model) isValidSize() bool {
	return m.size.Width >= m.minSize.Width && m.size.Height >= m.minSize.Height
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmd := m.resize(Size{
			Width:  msg.Width,
			Height: msg.Height,
		})

		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.Back) && !m.state.Focused():
			return m, m.back(1)
		case key.Matches(msg, m.keyMap.ShowHelp):
			m.help.ShowAll = !m.help.ShowAll
			cmd := m.setStateSize()
			return m, cmd
		}
	case notificationMsg:
		cmd := m.notify(msg.Message, m.notificationDefaultDuration)
		return m, cmd
	case notificationWithDurationMsg:
		cmd := m.notify(msg.Message, msg.Duration)
		return m, cmd
	case notificationTimeoutMsg:
		m.hideNotification()
		return m, nil
	case backMsg:
		return m, m.back(msg.Steps)
	case backToRootMsg:
		return m, m.back(m.history.Size())
	case pushStateMsg:
		return m, m.pushState(msg.State)
	case spinnerTickMsg:
		if !m.showSpinner {
			return m, nil
		}

		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)

		return m, func() tea.Msg {
			return spinnerTickMsg(cmd().(spinner.TickMsg))
		}
	case error:
		if errors.Is(msg, context.Canceled) {
			return m, nil
		}

		return m, m.onError(msg)
	}

	cmd := m.state.Update(m.ctx, msg)
	return m, cmd
}

func (m *Model) stateSize() Size {
	header := m.viewHeader()
	footer := m.viewFooter()

	size := m.size
	size.Height -= lipgloss.Height(header) + lipgloss.Height(footer)

	return size
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

func (m *Model) hideNotification() {
	m.notification.Message = ""
	if m.notification.Timer != nil {
		m.notification.Timer.Stop()
	}
}

func (m *Model) notify(message string, duration time.Duration) tea.Cmd {
	m.notification.Message = message

	if m.notification.Timer != nil {
		m.notification.Timer.Stop()
	}

	m.notification.Timer = time.NewTimer(duration)

	return func() tea.Msg {
		<-m.notification.Timer.C
		return notificationTimeoutMsg{}
	}
}
