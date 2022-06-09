package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

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
			newSelection := m.choices[m.selected]
			m.mediaURLs <- newSelection.StreamURL()
			currentSong := newSelection.CurrentTrack()
			m.notifier.Update(
				currentSong.Title,
				currentSong.Artist,
				"",
			)
		}

	case songMsg:
		smsg := songMsg(msg)
		var newChoices = make(models.Stations, len(m.choices))
		for i, choice := range m.choices {
			newChoice := choice
			if choice.Name() == smsg.StationName && choice.CurrentTrack().Title != smsg.Song.Title {
				newChoice.SetSong(smsg.Song)
				// Only send notification on selected station
				if i == m.selected {
					m.notifier.Update(
						smsg.Song.Title,
						smsg.Song.Artist,
						"",
					)
				}
			}
			newChoices[i] = newChoice
		}

		m.choices = newChoices
		return m, tea.Batch(waitForUpdates(m.updates), m.spinner.Tick)

	case errMsg:
		m.errMsg = errMsg(msg).err.Error()
		return m, tea.Batch(waitForUpdates(m.updates), m.spinner.Tick)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, m.spinner.Tick
}
