package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/dpark"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/sorcer"
	"github.com/codegoalie/bubbletea-test/wdwnt"
)

type tickMsg time.Time

type songMsg struct {
	Song        *models.TrackInfo
	StationName string
}

type errMsg struct {
	err error
}

func (e errMsg) Error() string { return e.err.Error() }

type model struct {
	choices  models.Stations
	cursor   int
	selected int

	lastTick time.Time
	errMsg   string
}

var initialModel = model{
	choices: models.Stations{
		models.NewRemoteStation(&sorcer.Atmospheres{}),
		models.NewRemoteStation(&dpark.Background{}),
		models.NewRemoteStation(&sorcer.Seasons{}),
		models.NewRemoteStation(&dpark.Christmas{}),
		models.NewRemoteStation(&sorcer.Mocha{}),
		models.NewRemoteStation(&wdwnt.Tunes{}),
		models.NewRemoteStation(&sorcer.Main{}),
		models.NewRemoteStation(&sorcer.SpaDay{}),
		models.NewRemoteStation(&dpark.Resort{}),
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
