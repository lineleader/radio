package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/codegoalie/bubbletea-test/models"
)

const maxSongLength = 75

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Station list\n\n")

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == i {
			checked = "x"
		}

		s.WriteString(cursor)
		s.WriteString(" [")
		s.WriteString(checked)
		s.WriteString("] ")
		s.WriteString(choice.Name())
		s.WriteString("\t")
		s.WriteString(truncate(choice.CurrentTrack().String(), maxSongLength))
		s.WriteString("\t")

		s.WriteString(remainingTimeDisplay(choice, m))

		s.WriteString("\n")
	}

	s.WriteString("\nPress q to quit.\n")

	return s.String()
}

func displayTime(left time.Duration) string {
	return fmt.Sprintf(
		"%02.f:%02.f",
		math.Floor(left.Minutes()),
		math.Mod(left.Seconds(), 60),
	)

}

func truncate(in string, maxLength int) string {
	if len(in) > maxLength {
		return in[:maxLength]
	}

	return fmt.Sprintf("%-*s", maxLength, in)
}

func remaining(currentTrack models.TrackInfo, now time.Time) time.Duration {
	return currentTrack.StartedAt.Add(currentTrack.Duration).Sub(now)
}

func remainingTimeDisplay(choice models.Station, m model) string {
	if choice.CurrentTrack().HideTiming {
		return "~"
	}

	s := strings.Builder{}

	remainingTime := remaining(choice.CurrentTrack(), m.lastTick)
	if remainingTime < 0 {
		s.WriteString(m.spinner.View())
	} else {
		s.WriteString("(")
		s.WriteString(displayTime(remainingTime))
		s.WriteString(" / ")
		s.WriteString(displayTime(choice.Duration()))
		s.WriteString(")")
	}

	return s.String()
}
