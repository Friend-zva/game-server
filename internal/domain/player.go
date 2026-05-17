package domain

import (
	"time"
)

type Player struct {
	Id           int
	State        StatePlayer
	CurrentFloor int
	Health       int
	BossDefeated bool

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
		CurrentFloor:     1,
		Health:           100,
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
