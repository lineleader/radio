package main

import (
	"fmt"
	"math"
	"time"
)

func (m model) View() string {
	s := "Station list\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == i {
			checked = "x"
		}

		s += fmt.Sprintf(
			"%s [%s] %s\t%s\t",
			cursor,
			checked,
			choice.Name(),
			choice.CurrentTrack(),
		)

		s += fmt.Sprintf(
			"(%s / %s)\n",
			displayTime(choice.Remaining(m.lastTick)),
			displayTime(choice.Duration()),
		)
	}

	s += "\nPress q to quit.\n"

	return s
}

func displayTime(left time.Duration) string {
	return fmt.Sprintf(
		"%02.f:%02.f",
		math.Floor(left.Minutes()),
		math.Mod(left.Seconds(), 60),
	)

}
