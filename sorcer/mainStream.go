package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

type Main struct{}

func (m Main) Name() string {
	return "Main Stream (Sorcer Radio)"
}

func (m Main) StreamURL() string {
	return "https://streaming.live365.com/a89268"
}

func (m Main) InfoURL() string {
	return "https://api.live365.com/station/a89268"
}

func (m Main) ParseTrackInfo(raw []byte) (models.TrackInfo, error) {
	return parseLive365TrackInfo(raw)
}

func (m Main) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(m, updates)
}
