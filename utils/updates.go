package utils

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
)

func SetupUpdateRegister(s models.PollingStation, updates chan models.TrackUpdate) tea.Cmd {
	return func() tea.Msg {
		ticks := 0
		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				ticks++
				if ticks%5 == 0 {
					update := models.TrackUpdate{
						StationName: s.Name(),
					}
					raw, err := HTTPGet(s.InfoURL())
					if err != nil {
						update.Error = fmt.Errorf("failed to get atmospheres info: %w", err)
					} else {
						update.Info, err = s.ParseTrackInfo(raw.Bytes())
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
