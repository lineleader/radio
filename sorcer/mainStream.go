package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const (
	// mainInfoURL = "https://api.live365.com/station/a89268"
	mainInfoURL = "https://listen.samcloud.com/webapi/station/67046/history?token=4e2d422c81d81eff066a193572925fa52962dd32&top=6&mediaTypeCodes=MUS,COM,NWS,INT&format=json&_=1662557064462"
)

type Main struct{}

func (m Main) Name() string {
	return "Main Stream (Sorcer Radio)"
}

func (m Main) StreamURL() string {
	// return "https://streaming.live365.com/a89268"
	return "https://samcloud.spacial.com/api/listen?sid=67046&rid=155533&f=mp3,any&br=128000,any&m=sc&t=ssl"
}

func (m Main) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		m.Name(),
		mainInfoURL,
		parseTrackInfo,
		updates,
	)
}
