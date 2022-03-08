package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
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

	case errMsg:
		m.errMsg = errMsg(msg).err.Error()
	}

	return m, nil
}

func tick() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func sync(station models.Station) tea.Cmd {
	trackInfo, err := latestSong(station)
	if err != nil {
		return func() tea.Msg {
			return errMsg{err}
		}
	}

	return func() tea.Msg {
		return songMsg{
			Song:        trackInfo,
			StationName: station.Name(),
		}
	}
}

func latestSong(station models.Station) (models.TrackInfo, error) {
	buf, err := utils.HTTPGet(station.InfoURL())
	if err != nil {
		err = fmt.Errorf("failed to get station info (%s): %w", station.Name(), err)
		return models.TrackInfo{}, err
	}

	if len(buf.Bytes()) == 0 {
		return models.TrackInfo{}, nil
	}

	return station.ParseTrackInfo(buf.Bytes())
}
