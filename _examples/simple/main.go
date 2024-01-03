package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda"
	"log"
)

func run() error {
	model := soda.New(New(), soda.WithMinSize(soda.Size{
		Width:  10,
		Height: 10,
	}))

	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
