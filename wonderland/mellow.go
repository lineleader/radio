package wonderland

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

const (
	mellowInfoURL = "wss://ctrl.radio-connected.co.uk/api/live/nowplaying/radio_wonderland_3"
)

type Mellow struct{}

func (m Mellow) Name() string {
	return "Mellow (Radio Wonderland)\t"
}

func (m Mellow) StreamURL() string {
	return "https://ctrl.radio-connected.co.uk/radio/8060/radio.mp3"
}

func (m Mellow) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return registerForUpdates(
		m.Name(),
		mellowInfoURL,
		updates,
	)
}
