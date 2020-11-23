package main

import "time"

type Station interface {
	Name() string
	CurrentTrack() string
	Remaining(time.Time) time.Duration
	Duration() time.Duration
}

type Stations []Station

type tickMsg time.Time

type model struct {
	choices  Stations
	cursor   int
	selected int

	lastTick time.Time
}
