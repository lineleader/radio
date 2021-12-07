package sorcer

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/codegoalie/bubbletea-test/models"
)

const (
	stationURL     = "http://listen.samcloud.com/webapi/station/%s/history"
	resultCount    = "5"
	mediaTypeCodes = "MUS,COM,NWS,INT"
	format         = "json"
)

type sorcerRadioSong struct {
	Title       string `json:"Title"`
	Album       string `json:"Album"`
	Artist      string `json:"Artist"`
	Duration    string `json:"Duration"`
	DatePlayed  string `json:"DatePlayed"`
	MediaItemID string `json:"MediaItemId"`
}

func infoURL(stationID, token string) string {
	uri, _ := url.Parse(fmt.Sprintf(stationURL, stationID))
	query := uri.Query()
	query.Add("token", token)
	query.Add("top", resultCount)
	query.Add("mediaTypeCodes", mediaTypeCodes)
	query.Add("format", format)
	query.Add("_", strconv.FormatInt(time.Now().Unix(), 10))
	uri.RawQuery = query.Encode()

	return uri.String()
}

func parseTrackInfo(raw []byte) (*models.TrackInfo, error) {
	recentSongs, err := unmarshalRecentSongs(raw)
	if err != nil {
		return nil, err
	}

	if len(recentSongs) < 1 {
		return &models.TrackInfo{}, nil
	}

	return recentToInfo(recentSongs[0])
}

func parseLive365TrackInfo(raw []byte) (*models.TrackInfo, error) {
	resp := live365Response{}
	err := json.Unmarshal(raw, &resp)
	if err != nil {
		var printable bytes.Buffer
		json.Indent(&printable, raw, "", "  ")
		err = fmt.Errorf(
			"failed to unmarshal sorcer radio history: %w\n%s",
			err,
			printable.String(),
		)
		return nil, err
	}

	track := resp.CurrentTrack
	info := models.TrackInfo{
		Title:     track.Title,
		Artist:    track.Artist,
		Duration:  time.Duration(track.Duration),
		StartedAt: time.Time(track.StartedAt),
	}

	return &info, nil

}

type live365Response struct {
	CurrentTrack live365Song `json:"current-track"`
}

type live365Song struct {
	Title     string          `json:"title"`
	Artist    string          `json:"artist"`
	Duration  live365Duration `json:"duration"`
	StartedAt live365Time     `json:"start"`

	Art     string      `json:"art"`
	EndedAt live365Time `json:"end"`
	// SyncOffset string      `json:"sync_offset"`
}

type live365Time time.Time

func (t *live365Time) UnmarshalJSON(b []byte) error {
	conv := strings.Replace(string(b), " ", "T", 1)
	conv = strings.Trim(conv, `"`)
	parsed, err := time.Parse(time.RFC3339Nano, conv)

	if err != nil {
		err = fmt.Errorf("failed to parse live365Time: %w", err)
		return err
	}

	*t = live365Time(parsed)
	return nil
}

type live365Duration time.Duration

func (d *live365Duration) UnmarshalJSON(b []byte) error {
	var duration string
	if string(b[0]) == "" {
		duration = strings.Trim(string(b), `"`) + "s"
	} else {
		bits := binary.LittleEndian.Uint32(b)
		float := math.Float64frombits(uint64(bits))
		duration = strconv.FormatFloat(float, 'f', -1, 64) + "s"
	}

	parsed, err := time.ParseDuration(duration)
	if err != nil {
		err = fmt.Errorf("failed to parse duration (%s): %w", string(b), err)
		return err
	}

	*d = live365Duration(parsed)
	return nil
}

func unmarshalRecentSongs(raw []byte) ([]sorcerRadioSong, error) {
	recentSongs := []sorcerRadioSong{}
	err := json.Unmarshal(raw, &recentSongs)
	if err != nil {
		err = fmt.Errorf(
			"failed to unmarshal sorcer radio history: %w (%s)",
			err,
			string(raw),
		)
		return nil, err
	}

	return recentSongs, nil
}

func recentToInfo(currentSong sorcerRadioSong) (*models.TrackInfo, error) {
	info := models.TrackInfo{}
	info.Title = currentSong.Title
	info.Artist = currentSong.Artist
	info.Album = currentSong.Album

	var err error
	info.Duration, err = time.ParseDuration(
		strings.TrimLeft(
			strings.ToLower(
				currentSong.Duration,
			),
			"pt",
		),
	)
	if err != nil {
		err = fmt.Errorf("failed to parse duration: %w", err)
		return nil, err
	}

	unixStr := strings.Split(strings.Trim(currentSong.DatePlayed, "\\/Date()"), "+")[0]
	unixMillisecs, err := strconv.ParseInt(unixStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse Sorcer started at: %w", err)
		return &info, err
	}
	startedAt := time.Unix(unixMillisecs/1000, 0)
	info.StartedAt = startedAt

	return &info, nil
}
