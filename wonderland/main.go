package wonderland

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

const (
	mainInfoURL = "wss://ctrl.radio-connected.co.uk/api/live/nowplaying/radio_wonderland_uk"
)

type Main struct{}

func (m Main) Name() string {
	return "Main (Radio Wonderland)\t"
}

func (m Main) StreamURL() string {
	return "https://ctrl.radio-connected.co.uk/radio/8040/radio.mp3"
}

func (m Main) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return registerForUpdates(
		m.Name(),
		mainInfoURL,
		updates,
	)
}
