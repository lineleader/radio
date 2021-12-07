package sorcer

import (
	"regexp"

	"github.com/codegoalie/bubbletea-test/models"
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

// InfoURL is the URL to fetch track data
func (s Atmospheres) InfoURL() string {
	return infoURL("130157", "acce5d6b010ebf1438bc1990f4cd357556aecf3b")
}

// ParseTrackInfo parses the provided bytes into a TrackInfo
func (s Atmospheres) ParseTrackInfo(raw []byte) (*models.TrackInfo, error) {
	recentSongs, err := unmarshalRecentSongs(raw)
	if err != nil {
		return nil, err
	}

	if len(recentSongs) < 1 {
		return &models.TrackInfo{}, nil
	}

	if len(recentSongs) > 1 && bigBandRegexp.MatchString(recentSongs[0].Title) {
		return recentToInfo(recentSongs[1])
	}

	return recentToInfo(recentSongs[0])
}
