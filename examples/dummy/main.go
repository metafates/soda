package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda"
)

func main() {
	program := tea.NewProgram(soda.New(New(), soda.WithAutoSize()), tea.WithAltScreen())

	_, err := program.Run()
	if err != nil {
		log.Fatal(err)
	}
}
