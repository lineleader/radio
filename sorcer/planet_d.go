package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const planetDName = "Atmospheres (Sorcer Radio)"
const plandeDStreamURL = "https://s43.radiolize.com/radio/8040/radio.mp3"

type PlanetD struct{}

// Name is the user presentable name for the stream
func (s PlanetD) Name() string {
	return atmospheresName
}

// StreamURL provides the current URL to stream audio
func (s PlanetD) StreamURL() string {
	return atmospheresStreamURL
}

func (s PlanetD) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		s.Name(),
		infoURL("130157", "acce5d6b010ebf1438bc1990f4cd357556aecf3b"),
		parseTrackInfoWithSkip(*bigBandRegexp),
		updates,
	)
}
