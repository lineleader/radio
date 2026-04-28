package sorcer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const ropeDropName = "Rope Drop (Sorcer Radio)\t"

// TODO the numeric subdomain is dynamic
// The web requests this URL to get the subdomain in the Servers part of the LiveStreamConfig struct below:
// https://playerservices.streamtheworld.com/api/livestream?mount=SP_R4852369&transports=http,hls&version=1.10&request.preventCache=1730918512938
// request.preventCache look liks a unix timestamp
// The sbmid in the ropeDropURL looks like a random UUID
const (
	ropeDropURL       = "https://19003.live.streamtheworld.com/SP_R3956488.mp3?tdsdk=js-2.9&swm=false&pname=TDSdk&pversion=2.9&banners=none&burst-time=15&sbmid=4ed46359-7d7e-4c38-aa41-15313d10dd5f"
	ropeDropStationID = "130153"
)

type RopeDrop struct{}

func (l RopeDrop) Name() string {
	return ropeDropName
}

func (l RopeDrop) StreamURL() string {
	return ropeDropURL
}

func (l RopeDrop) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		l.Name(),
		infoURL(ropeDropStationID, "15006229294361e8c7b1cfdb3ddc1bf9541527f4"),
		parseTrackInfo,
		updates,
	)
}
