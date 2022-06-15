package sorcer

import (
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const atmospheresName = "Atmospheres (Sorcer Radio)"
const atmospheresStreamURL = "https://samcloud.spacial.com/api/listen?sid=130157&m=sc&rid=273285"

var bigBandRegexp = regexp.MustCompile(`Friend Like Me \(Big Band\)`)

type Atmospheres struct{}

// Name is the user presentable name for the stream
func (s Atmospheres) Name() string {
	return atmospheresName
}

// StreamURL provides the current URL to stream audio
func (s Atmospheres) StreamURL() string {
	return atmospheresStreamURL
}

func (s Atmospheres) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {

	return utils.SetupUpdateRegister(
		s.Name(),
		infoURL("130157", "acce5d6b010ebf1438bc1990f4cd357556aecf3b"),
		parseTrackInfoWithSkip(*bigBandRegexp),
		updates,
	)
}
