package wdwnt

import (
	"encoding/json"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const tunesName = "WDWNTunes\t\t\t"
const tunesStreamURL = "https://streaming.live365.com/a31769"
const tunesInfoURL = "https://api.live365.com/station/a31769"

type Tunes struct{}

// Name is the userpresentable name of the stream
func (t Tunes) Name() string {
	return tunesName
}

// StreamURL provides the current URL to stream audio
func (t Tunes) StreamURL() string {
	return tunesStreamURL
}

func (t Tunes) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		t.Name(),
		tunesInfoURL,
		parseTrackInfo,
		updates,
	)
}

func parseTrackInfo(raw []byte) (models.TrackInfo, error) {
	resp := &wdwnTunesResponse{}
	err := json.Unmarshal(raw, &resp)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal WDWNTunes info: %w", err)
		return models.TrackInfo{}, err
	}
	info := models.TrackInfo{HideTiming: true}

	startedAt, err := time.Parse("2006-01-02 15:04:05-07:00", resp.CurrentTrack.Start)
	if err != nil {
		err = fmt.Errorf("failed to parse WDWNTunes started at info: %w", err)
		startedAt = time.Time{}
	}

	info.Title = resp.CurrentTrack.Title
	info.Artist = resp.CurrentTrack.Artist
	info.Album = ""
	info.Duration = time.Duration(resp.CurrentTrack.Duration * 1_000_000)
	info.StartedAt = startedAt

	return info, nil
}

type wdwnTunesResponse struct {
	CurrentTrack struct {
		Title    string  `json:"title"`
		Artist   string  `json:"artist"`
		Start    string  `json:"start"`
		Duration float64 `json:"duration"`
	} `json:"current-track"`
}
