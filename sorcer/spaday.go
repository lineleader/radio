package sorcer

import (
	"github.com/codegoalie/bubbletea-test/models"
)

const streamName = "Spa Day (Sorcer Radio)\t"
const spaStreamURL = "https://samcloud.spacial.com/api/listen?sid=130151&m=sc&rid=273274"

type SpaDay struct{}

func (s SpaDay) Name() string {
	return streamName
}

func (s SpaDay) StreamURL() string {
	return spaStreamURL
}

func (s SpaDay) InfoURL() string {
	return infoURL("130151", "29f4cfbac856cb4725f30257e21705772b59676d")
}

func (s SpaDay) ParseTrackInfo(raw []byte) (models.TrackInfo, error) {
	return parseTrackInfo(raw)
}
