package domain

import (
	"time"
)

type Player struct {
	Id           int
	State        StatePlayer
	FloorCurrent int
	Health       int

	EnteredDungeon bool
	BossDefeated   bool

	MonstersKilled map[int]int
	FloorsCleared  map[int]bool

	TimeEnterDungeon time.Time
	TimeLeaveDungeon time.Time
	TimeFloorEnter   map[int]time.Time
	TimeFloorClear   map[int]time.Duration
	TimeBossKill     time.Duration
}

func NewPlayer(id int, timeEnter time.Time) *Player {
	return &Player{
		Id:               id,
		State:            StatePlayerPlaying,
		FloorCurrent:     1,
		Health:           HealthMax,
		EnteredDungeon:   false,
		BossDefeated:     false,
		MonstersKilled:   make(map[int]int),
		FloorsCleared:    make(map[int]bool),
		TimeEnterDungeon: timeEnter,
		TimeLeaveDungeon: time.Time{},
		TimeFloorEnter:   make(map[int]time.Time),
		TimeFloorClear:   make(map[int]time.Duration),
		TimeBossKill:     0,
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
