package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const seasonsName = "Seasons (Sorcer Radio)\t"
const seasonsStreamURL = "https://samcloud.spacial.com/api/listen?sid=104853&rid=182288&f=mp3,any&br=128000,any&m=sc"

type Seasons struct{}

// Name is the user presentable name for the stream
func (s Seasons) Name() string {
	return seasonsName
}

// StreamURL provides the current URL to stream audio
func (s Seasons) StreamURL() string {
	return seasonsStreamURL
}

// InfoURL is the URL to fetch track data
func (s Seasons) InfoURL() string {
	return infoURL("104853", "254aec990e7d964645bc5fb68c58d45448f7719d")
}

// ParseTrackInfo parses the provided bytes into a TrackInfo
func (s Seasons) ParseTrackInfo(raw []byte) (models.TrackInfo, error) {
	return parseTrackInfo(raw)
}

func (s Seasons) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(s, updates)
}
