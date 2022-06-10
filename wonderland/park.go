package wonderland

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

const (
	parkInfoURL = "wss://ctrl.radio-connected.co.uk/api/live/nowplaying/radio_wonderland_2"
)

type Park struct{}

func (p Park) Name() string {
	return "Theme Park (Radio Wonderland)"
}

func (p Park) StreamURL() string {
	return "https://ctrl.radio-connected.co.uk/radio/8050/radio.mp3"
}

func (p Park) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return registerForUpdates(
		p.Name(),
		parkInfoURL,
		updates,
	)
}
