package dpark

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

// const christmasName = "Christmas (DPark Radio)\t"
const christmasName = "Halloween (DPark Radio)\t"
const christmasStreamURL = "https://listen.openstream.co/4287/;?1631785016772"
const christmasInfoURL = "https://c11.radioboss.fm/w/nowplayinginfo?u=39"

// Christmas streams the christmas music channel from DPark Radio
type Christmas struct{}

// Name is the userpresentable name of the stream
func (b Christmas) Name() string {
	return christmasName
}

// StreamURL provides the current URL to stream audio
func (b Christmas) StreamURL() string {
	return christmasStreamURL
}

func (b Christmas) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		b.Name(),
		christmasInfoURL,
		parseTrackInfo,
		updates,
	)
}
