package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Station interface {
	Name() string
	CurrentTrack() string
}

type staticStation struct {
	name         string
	currentTrack string
}

func (s staticStation) Name() string {
	return s.name
}

func (s staticStation) CurrentTrack() string {
	return s.currentTrack
}

type Stations []Station

type tickMsg float64

var programStartedAt = time.Now()

type model struct {
	choices  Stations
	cursor   int
	selected int

	lastTick float64
}

func (m model) Init() tea.Cmd {
	return tick()
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

	case tickMsg:
		m.lastTick = float64(msg)
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("Station list (%f)\n\n", m.lastTick)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == i {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name())
	}

	s += "\nPress q to quit.\n"

	return s
}

var initialModel = model{
	choices: Stations{
		staticStation{name: "DPark"},
		staticStation{name: "Sorcerer"},
		staticStation{name: "WDWNT"},
	},
}

func main() {
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting streamer: %v", err)
		os.Exit(1)
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t.Sub(programStartedAt).Seconds())
	})
}
