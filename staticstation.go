package main

import (
	"time"

	"github.com/codegoalie/bubbletea-test/models"
)

type staticStation struct {
	name string
	song *models.TrackInfo
}

func (s *staticStation) Name() string {
	return s.name
}

func (s *staticStation) InfoURL() string {
	return "https://lineleader.io"
}

func (s *staticStation) CurrentTrack() string {
	return s.song.String()
}

func (s *staticStation) Remaining(now time.Time) time.Duration {
	return s.song.StartedAt.Add(s.song.Duration).Sub(now)
}

func (s *staticStation) Duration() time.Duration {
	return s.song.Duration
}

func (s *staticStation) SetSong(newSong *models.TrackInfo) {
	s.song = newSong
}

func (s *staticStation) ParseTrackInfo(raw []byte) (*models.TrackInfo, error) {
	return &models.TrackInfo{
		Title:     "New Song",
		Album:     "New album",
		Artist:    "Singer",
		Duration:  50,
		StartedAt: time.Now(),
	}, nil
}
