package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Station struct {
	Name string
}

type Stations []Station

type model struct {
	choices  Stations
	cursor   int
	selected int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.cursor
		}

	}

	return m, nil
}

func (m model) View() string {
	s := "Station list\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == i {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}

	s += "\nPress q to quit.\n"

	return s
}

var initialModel = model{
	choices: Stations{
		{Name: "DPark"},
		{Name: "Sorcerer"},
		{Name: "WDWNT"},
	},
}

func main() {
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting streamer: %v", err)
		os.Exit(1)
	}
}
