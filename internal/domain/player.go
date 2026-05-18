package domain

import (
	"errors"
	"time"
)

var (
	ErrImpossibleMove    = errors.New("impossible move")
	ErrPlayerNotRegister = errors.New("player not registered")
	ErrPlayerDisqual     = errors.New("player disqualified")
	ErrPlayerNotEnter    = errors.New("player not enter dungeon")
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
	floors := make([]FloorMonsters, countFloors-1)

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

func (p *Player) EnterDungeon(timeEvent time.Time, timeOpened time.Time) error {
	if !p.TimeEnterDungeon.IsZero() || timeEvent.Before(timeOpened) {
		return ErrImpossibleMove
	}

	p.TimeEnterDungeon = timeEvent
	p.FloorsMonsters[p.FloorCurrent-1].Floor.TimeEnter = timeEvent
	return nil
}

func (p *Player) LeaveDungeon(timeEvent time.Time) error {
	if p.TimeEnterDungeon.IsZero() {
		return ErrImpossibleMove
	}

	p.TimeLeaveDungeon = timeEvent

	if p.FloorBoss.Cleared {
		p.State = StatePlayerSuccess
	} else {
		p.State = StatePlayerFail
	}

	return nil
}

func (p *Player) NextFloor(timeEvent time.Time, floorsMax int) error {
	if p.FloorCurrent >= floorsMax {
		return ErrImpossibleMove
	}

	p.FloorCurrent++

	var floor *Floor
	if p.FloorCurrent != floorsMax {
		floor = &(p.FloorsMonsters[p.FloorCurrent-1].Floor)
	} else {
		floor = &(p.FloorBoss)
	}

	if floor.TimeEnter.IsZero() {
		floor.TimeEnter = timeEvent
	}

	return nil
}

func (p *Player) PrevFloor() error {
	if p.FloorCurrent <= 1 {
		return ErrImpossibleMove
	}
	p.FloorCurrent--
	return nil
}

func (p *Player) EnterBossFloor(timeEvent time.Time, floorsMax int) error {
	if p.FloorCurrent != floorsMax {
		return ErrImpossibleMove
	}
	p.FloorBoss.TimeEnter = timeEvent
	return nil
}

func (p *Player) KillMonster(
	timeEvent time.Time, countMonstersPerFloor int, floorsMax int,
) error {
	if p.FloorCurrent == floorsMax {
		return ErrImpossibleMove
	}

	floor := &p.FloorsMonsters[p.FloorCurrent-1]
	if floor.MonstersKilled >= countMonstersPerFloor || floor.Floor.Cleared {
		return ErrImpossibleMove
	}

	floor.MonstersKilled++
	if floor.MonstersKilled == countMonstersPerFloor {
		floor.Floor.TimeClear = timeEvent
		floor.Floor.Cleared = true
	}

	return nil
}

func (p *Player) KillBoss(timeEvent time.Time, floorsMax int) error {
	if p.FloorBoss.Cleared || p.FloorCurrent != floorsMax {
		return ErrImpossibleMove
	}

	p.FloorBoss.Cleared = true
	p.FloorBoss.TimeClear = timeEvent
	return nil
}

func (p *Player) CannotContinue(timeEvent time.Time) {
	p.State = StatePlayerDisqual
	p.TimeLeaveDungeon = timeEvent
}

func (p *Player) RestoreHealth(amount int) {
	p.Health += amount

	if p.Health > HealthMax {
		p.Health = HealthMax
	}
}

func (p *Player) ReceiveDamage(amount int, timeEvent time.Time) {
	p.Health -= amount

	if p.Health <= 0 {
		p.Health = 0
		p.State = StatePlayerFail
		p.TimeLeaveDungeon = timeEvent
	}
}
