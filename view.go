package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

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
		s.WriteString(choice.CurrentTrack().String())
		s.WriteString("\t")

		remainingTime := choice.Remaining(m.lastTick)
		if remainingTime < 0 {
			s.WriteString("(Loading...)")
		} else {
			s.WriteString("(")
			s.WriteString(displayTime(remainingTime))
			s.WriteString(" / ")
			s.WriteString(displayTime(choice.Duration()))
			s.WriteString(")")
		}

		s.WriteString("\n")
	}

	if m.errMsg != "" {
		s.WriteString("\n" + m.errMsg + "\n")
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
