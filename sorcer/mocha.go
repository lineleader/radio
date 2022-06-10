package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const mochaName = "Mocha (Sorcer Radio)\t"
const mochaStreamURL = "https://samcloud.spacial.com/api/listen?sid=100903&m=sc&rid=177361"

type Mocha struct{}

// Name is the user presentable name for the stream
func (s Mocha) Name() string {
	return mochaName
}

// StreamURL provides the current URL to stream audio
func (s Mocha) StreamURL() string {
	return mochaStreamURL
}

func (m Mocha) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		m.Name(),
		infoURL("100903", "030c8d06bdd9e82eae632eaff484df864c54f14c"),
		parseTrackInfo,
		updates,
	)
}
