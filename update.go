package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

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
		if m.lastTick.Unix()%5 != 0 {
			return m, tick()
		}

		cmds := []tea.Cmd{tick()}
		for _, choice := range m.choices {
			if choice.Remaining(time.Now()) < 0 {
				cmds = append(cmds, sync(choice))
			}
		}
		return m, tea.Batch(cmds...)

	case songMsg:
		smsg := songMsg(msg)
		var newChoices = make(models.Stations, len(m.choices))
		for i, choice := range m.choices {
			newChoice := choice
			if choice.Name() == smsg.StationName {
				newChoice.SetSong(smsg.Song)
			}
			newChoices[i] = newChoice
		}

		m.choices = newChoices
	}

	return m, nil
}

func tick() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func sync(station models.Station) tea.Cmd {
	return func() tea.Msg {
		return songMsg{
			Song: models.Song{
				Name:     "New song",
				Duration: 45 * time.Second,
				EndsAt:   time.Now().Add(32 * time.Second),
			},
			StationName: station.Name(),
		}
	}
}
