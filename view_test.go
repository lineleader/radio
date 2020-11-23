package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	m := model{}
	expected := "Station list"
	actual := m.View()

	assert.Contains(t, actual, expected)
}

func TestListStationNames(t *testing.T) {
	m := model{
		choices: Stations{
			staticStation{name: "Station0"},
			staticStation{name: "Station1"},
		},
	}
	actual := m.View()

	for i, choice := range m.choices {
		t.Run("Listing station"+strconv.Itoa(i), func(it *testing.T) {
			assert.Contains(it, actual, choice.Name())
		})
	}
}

func TestShowDuration(t *testing.T) {
	m := model{
		choices: Stations{
			staticStation{
				name:     "Test",
				duration: time.Minute,
			},
		},
	}

	actual := m.View()

	assert.Contains(t, actual, "/ 01:00)")
}

func TestShowRemaining(t *testing.T) {
	m := model{
		choices: Stations{
			staticStation{
				name:     "Test",
				duration: time.Minute,
				endsAt:   time.Now().Add(3 * time.Second),
			},
		},
		lastTick: time.Now(),
	}

	actual := m.View()

	assert.Contains(t, actual, "(00:03 /")
}

func TestShowLoading(t *testing.T) {
	m := model{
		choices: Stations{
			staticStation{
				name:     "Test",
				duration: time.Minute,
				endsAt:   time.Now().Add(-3 * time.Second),
			},
		},
		lastTick: time.Now(),
	}

	actual := m.View()

	assert.Contains(t, actual, "(Loading...)")
}
