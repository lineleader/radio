package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const (
	mainInfoURL = "https://api.live365.com/station/a89268"
)

type Main struct{}

func (m Main) Name() string {
	return "Main Stream (Sorcer Radio)"
}

func (m Main) StreamURL() string {
	return "https://streaming.live365.com/a89268"
}

func (m Main) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		m.Name(),
		mainInfoURL,
		parseLive365TrackInfo,
		updates,
	)
}
