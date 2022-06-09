package models

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type StationSource interface {
	Name() string
	StreamURL() string
	RegisterForUpdates(chan TrackUpdate) tea.Cmd

	// deprecated
	InfoURL() string
	ParseTrackInfo([]byte) (TrackInfo, error)
}

type RemoteStation struct {
	source       StationSource
	currentTrack TrackInfo
}

func NewRemoteStation(source StationSource) *RemoteStation {
	return &RemoteStation{source: source, currentTrack: TrackInfo{}}
}

func (r *RemoteStation) Name() string {
	return r.source.Name()
}

func (r *RemoteStation) CurrentTrack() TrackInfo {
	return r.currentTrack
}

func (r *RemoteStation) Remaining(now time.Time) time.Duration {
	return r.currentTrack.StartedAt.Add(r.currentTrack.Duration).Sub(now)
}

func (r *RemoteStation) Duration() time.Duration {
	return r.currentTrack.Duration
}

func (r *RemoteStation) SetSong(newTrack TrackInfo) {
	r.currentTrack = newTrack
}

func (r *RemoteStation) StreamURL() string {
	return r.source.StreamURL()
}

func (r *RemoteStation) RegisterForUpdates(updates chan TrackUpdate) tea.Cmd {
	return r.source.RegisterForUpdates(updates)
}
