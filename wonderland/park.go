package wonderland

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/gorilla/websocket"
)

type Park struct{}

func (p Park) Name() string {
	return "Theme Park Stream (Radio Wonderland)"
}

func (p Park) StreamURL() string {
	return "https://ctrl.radio-connected.co.uk/radio/8050/radio.mp3"
}

func (p Park) InfoURL() string {
	return "wss://ctrl.radio-connected.co.uk/api/live/nowplaying/radio_wonderland_2"
}

func (p Park) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return func() tea.Msg {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		c, _, err := websocket.DefaultDialer.Dial(p.InfoURL(), nil)
		if err != nil {
			log.Fatal("failed to dial:", err)
		}
		defer c.Close()

		done := make(chan struct{})

		go func() {
			defer close(done)
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("failed to read:", err)
					return
				}
				update := models.TrackUpdate{StationName: p.Name()}
				update.Info, err = p.ParseTrackInfo([]byte(message))
				if err != nil {
					update.Error = fmt.Errorf("failed to : %w", err)
				}
				updates <- update
			}
		}()

		for {
			select {
			case <-done:
				return tea.Quit
			case <-interrupt:
				log.Println("interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(
					websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
				)
				if err != nil {
					log.Println("failed to write close:", err)
					return tea.Quit
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return tea.Quit
			}
		}
	}
}

func (p Park) ParseTrackInfo(raw []byte) (models.TrackInfo, error) {
	var msg songMessage
	err := json.Unmarshal(raw, &msg)
	if err != nil {
		err = fmt.Errorf("failed to parse wonderland park raw track info: %w", err)
		return models.TrackInfo{}, err
	}

	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		err = fmt.Errorf("failed to load New York time zone: %w", err)
		return models.TrackInfo{}, err
	}

	song := msg.NowPlaying.Song
	return models.TrackInfo{
		Title:     song.Title,
		Album:     song.Album,
		Artist:    song.Artist,
		Duration:  time.Duration(time.Duration(msg.NowPlaying.Duration) * time.Second),
		StartedAt: time.Unix(msg.NowPlaying.PlayedAt, 0).In(est),
	}, nil
}

type songMessage struct {
	Station struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Shortcode       string `json:"shortcode"`
		Description     string `json:"description"`
		Frontend        string `json:"frontend"`
		Backend         string `json:"backend"`
		ListenURL       string `json:"listen_url"`
		URL             string `json:"url"`
		PublicPlayerURL string `json:"public_player_url"`
		PlaylistPlsURL  string `json:"playlist_pls_url"`
		PlaylistM3UURL  string `json:"playlist_m3u_url"`
		IsPublic        bool   `json:"is_public"`
		Mounts          []struct {
			Path      string `json:"path"`
			IsDefault bool   `json:"is_default"`
			ID        int    `json:"id"`
			Name      string `json:"name"`
			URL       string `json:"url"`
			Bitrate   int    `json:"bitrate"`
			Format    string `json:"format"`
			Listeners struct {
				Total   int `json:"total"`
				Unique  int `json:"unique"`
				Current int `json:"current"`
			} `json:"listeners"`
		} `json:"mounts"`
		Remotes []interface{} `json:"remotes"`
	} `json:"station"`
	Listeners struct {
		Total   int `json:"total"`
		Unique  int `json:"unique"`
		Current int `json:"current"`
	} `json:"listeners"`
	Live struct {
		IsLive         bool        `json:"is_live"`
		StreamerName   string      `json:"streamer_name"`
		BroadcastStart interface{} `json:"broadcast_start"`
	} `json:"live"`
	NowPlaying struct {
		Elapsed   int    `json:"elapsed"`
		Remaining int    `json:"remaining"`
		ShID      int    `json:"sh_id"`
		PlayedAt  int64  `json:"played_at"`
		Duration  int    `json:"duration"`
		Playlist  string `json:"playlist"`
		Streamer  string `json:"streamer"`
		IsRequest bool   `json:"is_request"`
		Song      struct {
			ID           string        `json:"id"`
			Text         string        `json:"text"`
			Artist       string        `json:"artist"`
			Title        string        `json:"title"`
			Album        string        `json:"album"`
			Genre        string        `json:"genre"`
			Lyrics       string        `json:"lyrics"`
			Art          string        `json:"art"`
			CustomFields []interface{} `json:"custom_fields"`
		} `json:"song"`
	} `json:"now_playing"`
	PlayingNext struct {
		CuedAt    int    `json:"cued_at"`
		Duration  int    `json:"duration"`
		Playlist  string `json:"playlist"`
		IsRequest bool   `json:"is_request"`
		Song      struct {
			ID           string        `json:"id"`
			Text         string        `json:"text"`
			Artist       string        `json:"artist"`
			Title        string        `json:"title"`
			Album        string        `json:"album"`
			Genre        string        `json:"genre"`
			Lyrics       string        `json:"lyrics"`
			Art          string        `json:"art"`
			CustomFields []interface{} `json:"custom_fields"`
		} `json:"song"`
	} `json:"playing_next"`
	SongHistory []struct {
		ShID      int    `json:"sh_id"`
		PlayedAt  int    `json:"played_at"`
		Duration  int    `json:"duration"`
		Playlist  string `json:"playlist"`
		Streamer  string `json:"streamer"`
		IsRequest bool   `json:"is_request"`
		Song      struct {
			ID           string        `json:"id"`
			Text         string        `json:"text"`
			Artist       string        `json:"artist"`
			Title        string        `json:"title"`
			Album        string        `json:"album"`
			Genre        string        `json:"genre"`
			Lyrics       string        `json:"lyrics"`
			Art          string        `json:"art"`
			CustomFields []interface{} `json:"custom_fields"`
		} `json:"song"`
	} `json:"song_history"`
	IsOnline bool   `json:"is_online"`
	Cache    string `json:"cache"`
}
