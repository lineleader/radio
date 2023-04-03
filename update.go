package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.updateKeyMsg(msg)
	case songMsg:
		return m.updateSongMsg(songMsg(msg))
	case errMsg:
		return m.updateErrMsg(errMsg(msg))
	case spinner.TickMsg:
		return m.updateSpinner(msg)
	case tickMsg:
	}
	return m, m.spinner.Tick
}

func (m model) updateKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		} else {
			m.cursor = len(m.choices) - 1
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		} else {
			m.cursor = 0
		}
	case "enter", " ":
		return m.updateEnter()
	}

	return m, m.spinner.Tick
}

func (m model) updateEnter() (tea.Model, tea.Cmd) {
	m.selected = m.cursor
	newSelection := m.choices[m.selected]
	m.mediaURLs <- newSelection.StreamURL()
	currentSong := newSelection.CurrentTrack()
	m.notifier.Update(
		currentSong.Title,
		currentSong.Artist,
		"",
	)
	return m, tea.Batch(waitForUpdates(m.updates), m.spinner.Tick)
}

func (m model) updateSongMsg(msg songMsg) (tea.Model, tea.Cmd) {
	var newChoices = make(models.Stations, len(m.choices))
	for i, choice := range m.choices {
		newChoice := choice
		if choice.Name() == msg.StationName && choice.CurrentTrack().Title != msg.Song.Title {
			newChoice.SetSong(msg.Song)
			// Only send notification on selected station
			if i == m.selected {
				m.notifier.Update(
					msg.Song.Title,
					msg.Song.Artist,
					"",
				)
			}
		}
		newChoices[i] = newChoice
	}
	m.choices = newChoices
	return m, tea.Batch(waitForUpdates(m.updates), m.spinner.Tick)
}

func (m model) updateErrMsg(msg errMsg) (tea.Model, tea.Cmd) {
	m.errMsg = msg.err.Error()
	return m, tea.Batch(waitForUpdates(m.updates), m.spinner.Tick)
}

func (m model) updateSpinner(msg spinner.TickMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	m.lastTick = spinner.TickMsg(msg).Time
	return m, cmd
}
