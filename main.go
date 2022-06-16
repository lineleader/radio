package main

import (
	"fmt"
	"log"
	"os"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codegoalie/bubbletea-test/dpark"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/sorcer"
	"github.com/codegoalie/bubbletea-test/wdwnt"
	"github.com/codegoalie/bubbletea-test/wonderland"
	"github.com/codegoalie/golibnotify"
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
	spinner   spinner.Model
	errMsg    string
	mediaURLs chan string
	actions   chan string
	notifier  *golibnotify.SimpleNotifier
	updates   chan models.TrackUpdate
}

func (m model) Init() tea.Cmd {
	cmds := []tea.Cmd{waitForUpdates(m.updates), m.spinner.Tick}
	for _, station := range m.choices {
		cmds = append(cmds, station.RegisterForUpdates(m.updates))
	}
	return tea.Batch(cmds...)
}

type updateMsg models.TrackUpdate

func waitForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return func() tea.Msg {
		update := <-updates
		if update.Error != nil {
			return errMsg{update.Error}
		}
		return songMsg{Song: update.Info, StationName: update.StationName}
	}
}

var initialModel = model{
	choices: models.Stations{
		models.NewRemoteStation(&sorcer.Atmospheres{}),
		models.NewRemoteStation(&dpark.Background{}),
		models.NewRemoteStation(&wonderland.Park{}),
		models.NewRemoteStation(&sorcer.Seasons{}),
		models.NewRemoteStation(&sorcer.Mocha{}),
		models.NewRemoteStation(&wdwnt.Tunes{}),
		models.NewRemoteStation(&sorcer.Main{}),
		models.NewRemoteStation(&wonderland.Main{}),
		models.NewRemoteStation(&sorcer.SpaDay{}),
		models.NewRemoteStation(&wonderland.Mellow{}),
		models.NewRemoteStation(&dpark.Christmas{}),
		models.NewRemoteStation(&dpark.Resort{}),
	},
	lastTick: time.Now(),

	mediaURLs: make(chan string),
	notifier:  golibnotify.NewSimpleNotifier("Stream Player"),
	updates:   make(chan models.TrackUpdate),
}

func main() {
	quit := make(chan struct{})
	defer close(quit)
	defer initialModel.notifier.Close()

	go playAudio(initialModel.mediaURLs, quit)
	initialModel.mediaURLs <- initialModel.choices[initialModel.selected].StreamURL()

	initialModel.spinner = spinner.New()
	initialModel.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	initialModel.spinner.Spinner = spinner.MiniDot

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
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
