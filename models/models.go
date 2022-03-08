package models

import (
	"strings"
	"time"
)

type Station interface {
	Name() string
	CurrentTrack() TrackInfo
	Remaining(time.Time) time.Duration
	Duration() time.Duration
	SetSong(TrackInfo)

	StreamURL() string
	InfoURL() string
	ParseTrackInfo(raw []byte) (TrackInfo, error)
}

type Stations []Station

type TrackInfo struct {
	Title     string
	Album     string
	Artist    string
	Duration  time.Duration
	StartedAt time.Time
}

func (t TrackInfo) String() string {
	var msg strings.Builder
	msg.WriteString(t.Title)

	if t.Artist != "" {
		msg.WriteString(" - ")
		msg.WriteString(t.Artist)
	}

	if t.Album != "" {
		msg.WriteString(" [")
		msg.WriteString(t.Album)
		msg.WriteString("]")
	}

	return msg.String()
}
