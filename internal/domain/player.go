package domain

import (
	"time"
)

type Player struct {
	Id           int
	State        StatePlayer
	FloorCurrent int
	Health       int

	FloorsMonsters []FloorMonsters
	FloorBoss      Floor

	TimeEnterDungeon time.Time
	TimeLeaveDungeon time.Time
}

func NewPlayer(id, countFloors int) *Player {
	floors := make([]FloorMonsters, countFloors)

	return &Player{
		Id:               id,
		State:            StatePlayerPlaying,
		FloorCurrent:     1,
		Health:           HealthMax,
		FloorsMonsters:   floors,
		FloorBoss:        Floor{},
		TimeEnterDungeon: time.Time{},
		TimeLeaveDungeon: time.Time{},
	}
}

func (p *Player) RestoreHealth(amount int) {
	p.Health += amount

	if p.Health > HealthMax {
		p.Health = HealthMax
	}
}

func (p *Player) ReceiveDamage(amount int) []EventOutgoingID {
	p.Health -= amount

	var events []EventOutgoingID

	if p.Health <= 0 {
		p.Health = 0
		p.State = StatePlayerFail
		events = append(events, EventOutgoingDead)
	}

	return events
}
