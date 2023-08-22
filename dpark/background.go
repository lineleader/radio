package dpark

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const (
	backgroundName      = "Background (DPark Radio)\t"
	backgroundStreamURL = "https://listen.openstream.co/3127/audio?"
	backgroundInfoURL   = "https://c5.radioboss.fm/w/nowplayinginfo?u=38&_="
)

// const backgroundStreamURL = "https://str2b.openstream.co/578?aw_0_1st.collectionid=3127&aw_0_1st.publisherId=602"

// Background streams the background music channel from DPark Radio
type Background struct{}

// Name is the userpresentable name of the stream
func (b Background) Name() string {
	return backgroundName
}

// StreamURL provides the current URL to stream audio
func (b Background) StreamURL() string {
	return backgroundStreamURL + fmt.Sprintf("%d", time.Now().Unix())
}

func (b Background) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		b.Name(),
		backgroundInfoURL+fmt.Sprintf("%d", time.Now().Unix()),
		parseTrackInfo,
		updates)
}
