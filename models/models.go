package models

import "time"

type Station interface {
	Name() string
	CurrentTrack() string
	Remaining(time.Time) time.Duration
	Duration() time.Duration
	Song() *Song
	SetSong(Song)
}

type Stations []Station

type Song struct {
	Name     string
	Duration time.Duration
	EndsAt   time.Time
}

type tickMsg time.Time

type songMsg struct {
	Song        Song
	StationName string
}

type model struct {
	choices  Stations
	cursor   int
	selected int

	lastTick time.Time
}

type TrackInfo struct {
	Title     string
	Album     string
	Artist    string
	Duration  float64
	StartedAt time.Time
}
