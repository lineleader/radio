package utils

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

type trackInfoParser func([]byte) (models.TrackInfo, error)

func SetupUpdateRegister(stationName, infoURL string, parseTrackInfo trackInfoParser, updates chan models.TrackUpdate) tea.Cmd {
	return func() tea.Msg {
		ticks := 0
		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				ticks++
				// TODO: Be smarter about when to update (like no more remaining)
				if ticks%5 == 0 {
					update := models.TrackUpdate{
						StationName: stationName,
					}
					raw, err := HTTPGet(infoURL)
					if err != nil {
						update.Error = fmt.Errorf("failed to get atmospheres info: %w", err)
					} else {
						update.Info, err = parseTrackInfo(raw.Bytes())
						if err != nil {
							update.Error = fmt.Errorf("failed to parse atmospheres info: %w", err)
						}
					}

					updates <- update
				}
			}
			// Probably need to handle quitting or other cases here
		}
	}
}
