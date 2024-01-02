package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda"
	"github.com/metafates/soda/title"
)

var _ soda.State = (*State)(nil)

type State struct {
	size   soda.Size
	keyMap keyMap
}

func (s *State) Status() string {
	return ""
}

func (s *State) Destroy() {
}

func (s *State) Backable() bool {
	return true
}

func (s *State) Resize(size soda.Size) tea.Cmd {
	s.size = size
	return nil
}

func (s *State) Title() title.Title {
	return title.New("Dummy")
}

func (s *State) Subtitle() string {
	return "Subtitle"
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(mh soda.ModelHandler) tea.Cmd {
	return nil
}

func (s *State) Update(mh soda.ModelHandler, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.SendNotification):
			return soda.NotifyTimeout(fmt.Sprint(time.Now().Unix()), time.Second)
		case key.Matches(msg, s.keyMap.RunTask):
			return soda.Do("Running task", func() tea.Msg {
				time.Sleep(time.Second * 5)
				return nil
			})
		}
	}
	return nil
}

func (s *State) View(mh soda.ModelHandler) string {
	return fmt.Sprintf("Width %d\nHeight %d", s.size.Width, s.size.Height)
}
