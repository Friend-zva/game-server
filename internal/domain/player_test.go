package domain

import (
	"testing"
	"time"
)

func TestPlayerNextFloor(t *testing.T) {
	tests := []struct {
		name         string
		floorCurrent int
		countFloors  int
		floorsMax    int
		errExp       error
	}{
		{
			name:         "Next Floor is possible",
			floorCurrent: 1,
			countFloors:  3,
			floorsMax:    3,
			errExp:       nil,
		},
		{
			name:         "Next Floor is impossible",
			floorCurrent: 3,
			countFloors:  3,
			floorsMax:    3,
			errExp:       ErrImpossibleMove,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeEvent := time.Now()

			p := NewPlayer(1, tt.countFloors)
			p.FloorCurrent = tt.floorCurrent

			err := p.NextFloor(timeEvent, tt.floorsMax)
			if err != tt.errExp {
				t.Errorf("expected error to %v, got %v", tt.errExp, err)
			}

			if err == nil && p.FloorCurrent != tt.floorCurrent+1 {
				t.Errorf("expected floor to %d, got %d",
					p.FloorCurrent,
					tt.floorCurrent+1,
				)
			}
		})
	}
}

func TestPlayerReceiveDamage(t *testing.T) {
	tests := []struct {
		name                string
		health              int
		state               StatePlayer
		damage              int
		healthExp           int
		stateExp            StatePlayer
		timeLeaveDungeonExp bool
	}{
		{
			name:                "Player survives damage",
			health:              50,
			state:               StatePlayerPlaying,
			damage:              20,
			healthExp:           30,
			stateExp:            StatePlayerPlaying,
			timeLeaveDungeonExp: false,
		},
		{
			name:                "Player dies from damage",
			health:              50,
			state:               StatePlayerPlaying,
			damage:              50,
			healthExp:           0,
			stateExp:            StatePlayerFail,
			timeLeaveDungeonExp: true,
		},
		{
			name:                "Health cannot go below zero",
			health:              10,
			state:               StatePlayerPlaying,
			damage:              999,
			healthExp:           0,
			stateExp:            StatePlayerFail,
			timeLeaveDungeonExp: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeEvent := time.Now()

			p := NewPlayer(1, 3)
			p.Health = tt.health
			p.State = tt.state

			p.ReceiveDamage(tt.damage, timeEvent)

			if p.Health != tt.healthExp {
				t.Errorf("expected health to be %d, got %d", tt.healthExp, p.Health)
			}

			if p.State != tt.stateExp {
				t.Errorf("expected state to be %s, got %s", tt.stateExp, p.State)
			}

			if tt.timeLeaveDungeonExp &&
				p.TimeLeaveDungeon != timeEvent {
				t.Errorf("expected TimeLeaveDungeon to be set")
			}
		})
	}
}
