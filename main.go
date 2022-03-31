package main

import (
	"fmt"
	"log"
	"os"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/dpark"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/sorcer"
	"github.com/codegoalie/bubbletea-test/wdwnt"
)

type tickMsg time.Time

type songMsg struct {
	Song        models.TrackInfo
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

	lastTick  time.Time
	errMsg    string
	mediaURLs chan string
	actions   chan string
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

	mediaURLs: make(chan string),
}

func main() {
	quit := make(chan struct{})
	defer close(quit)

	go playAudio(initialModel.mediaURLs, quit)
	initialModel.mediaURLs <- initialModel.choices[initialModel.selected].StreamURL()

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting streamer: %v", err)
		os.Exit(1)
	}
}

func playAudio(nextMediaURL <-chan string, quit chan struct{}) {
	// Initialize libvlc. Additional command line arguments can be passed in
	// to libvlc by specifying them in the Init function.
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal("failed to init vlc", err)
	}
	defer vlc.Release()

	// Create a new player.
	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal("failed to create new vlc player: ", err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Retrieve player event manager.
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal("failed to get vlc player event manager", err)
	}

	// Register the media end reached event with the event manager.
	eventCallback := func(event vlc.Event, userData interface{}) {
		close(quit)
	}

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal("failed to attach to media end reached event", err)
	}
	defer manager.Detach(eventID)

	var media *vlc.Media

play:
	for {
		select {
		case currentMediaURL := <-nextMediaURL:
			if media != nil {
				media.Release()
			}
			media, err = player.LoadMediaFromURL(currentMediaURL)
			if err != nil {
				log.Fatal("failed to load media from url", err)
			}

			// Start playing the media.
			err = player.Play()
			if err != nil {
				log.Fatal("failed to play media", err)
			}
		case <-quit:
			media.Release()
			break play
		}
	}
}
