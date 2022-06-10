package models

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Station interface {
	StationSource
	CurrentTrack() TrackInfo
	Duration() time.Duration
	SetSong(TrackInfo)
}

type StationSource interface {
	Name() string
	StreamURL() string
	RegisterForUpdates(chan TrackUpdate) tea.Cmd
}

type Stations []Station

type TrackUpdate struct {
	StationName string
	Info        TrackInfo
	Error       error
}

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
