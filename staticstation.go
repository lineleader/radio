package main

import (
	"time"

	"github.com/codegoalie/bubbletea-test/models"
)

type staticStation struct {
	name string
	song models.Song
}

func (s *staticStation) Name() string {
	return s.name
}

func (s *staticStation) CurrentTrack() string {
	return s.song.Name
}

func (s *staticStation) Duration() time.Duration {
	return s.song.Duration
}

func (s *staticStation) Remaining(now time.Time) time.Duration {
	return s.song.EndsAt.Sub(now)
}

func (s *staticStation) Song() *models.Song {
	return &s.song
}

func (s *staticStation) SetSong(newSong models.Song) {
	s.song = newSong
}
