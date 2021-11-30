package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

type tickMsg time.Time

type songMsg struct {
	Song        models.Song
	StationName string
}

type model struct {
	choices  models.Stations
	cursor   int
	selected int

	lastTick time.Time
}

var initialModel = model{
	choices: models.Stations{
		// &sorcer.Seasons{},
		&staticStation{
			name: "WDWNTunes",
			song: models.Song{
				EndsAt:   time.Now().Add(5 * time.Second),
				Duration: time.Minute * 2,
				Name:     "Comercial",
			},
		},
		&staticStation{
			name: "DPark",
			song: models.Song{
				EndsAt:   time.Now().Add(5 * time.Second),
				Duration: time.Minute * 2,
				Name:     "Christmas",
			},
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
