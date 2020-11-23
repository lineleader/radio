package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var initialModel = model{
	choices: Stations{
		staticStation{
			name:         "DPark",
			endsAt:       time.Now().Add(time.Minute),
			duration:     time.Minute * 2,
			currentTrack: "Christmas",
		},
		staticStation{
			name:         "Sorcerer",
			endsAt:       time.Now().Add(3 * time.Second),
			duration:     time.Minute*2 + time.Second*23,
			currentTrack: "Parade",
		},
		staticStation{
			name:         "WDWNT",
			endsAt:       time.Now().Add(time.Minute).Add(3 * time.Second),
			duration:     time.Minute*2 + time.Second*32,
			currentTrack: "EPCOT enterance",
		},
	},
	lastTick: time.Now(),
}

func main() {
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting streamer: %v", err)
		os.Exit(1)
	}
}
