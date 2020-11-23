package main

import (
	"fmt"
	"math"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Station interface {
	Name() string
	CurrentTrack(time.Time) string
}

type staticStation struct {
	name         string
	currentTrack string
	duration     time.Duration
	endsAt       time.Time
}

func (s staticStation) Name() string {
	return s.name
}

func (s staticStation) CurrentTrack(now time.Time) string {
	return fmt.Sprintf(
		"%s (%s / %s)",
		s.currentTrack,
		displayTime(s.endsAt.Sub(now)),
		displayTime(s.duration),
	)
}

func displayTime(left time.Duration) string {
	return fmt.Sprintf(
		"%02.f:%02.f",
		math.Floor(left.Minutes()),
		math.Mod(left.Seconds(), 60),
	)

}

type Stations []Station

type tickMsg time.Time

var programStartedAt = time.Now()

type model struct {
	choices  Stations
	cursor   int
	selected int

	lastTick time.Time
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
		m.lastTick = time.Time(msg)
		return m, tick()
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

		s += fmt.Sprintf(
			"%s [%s] %s\t%s\n",
			cursor,
			checked,
			choice.Name(),
			choice.CurrentTrack(m.lastTick),
		)
	}

	s += "\nPress q to quit.\n"

	return s
}

var initialModel = model{
	choices: Stations{
		staticStation{
			name:     "DPark",
			endsAt:   time.Now().Add(time.Minute),
			duration: time.Minute * 2,
		},
		staticStation{
			name:     "Sorcerer",
			endsAt:   time.Now().Add(time.Minute).Add(13 * time.Second),
			duration: time.Minute*2 + time.Second*23,
		},
		staticStation{
			name:     "WDWNT",
			endsAt:   time.Now().Add(time.Minute).Add(3 * time.Second),
			duration: time.Minute*2 + time.Second*32,
		},
	},
	lastTick: time.Now(),
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
		return tickMsg(t)
	})
}
