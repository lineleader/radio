package models

import "time"

type StationSource interface {
	Name() string
	StreamURL() string
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

func (r *RemoteStation) InfoURL() string {
	return r.source.InfoURL()
}

func (r *RemoteStation) ParseTrackInfo(raw []byte) (TrackInfo, error) {
	return r.source.ParseTrackInfo(raw)
}
