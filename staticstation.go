package main

import "time"

type staticStation struct {
	name         string
	currentTrack string
	duration     time.Duration
	endsAt       time.Time
}

func (s staticStation) Name() string {
	return s.name
}

func (s staticStation) CurrentTrack() string {
	return s.currentTrack
}

func (s staticStation) Duration() time.Duration {
	return s.duration
}

func (s staticStation) Remaining(now time.Time) time.Duration {
	return s.endsAt.Sub(now)
}
