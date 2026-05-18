package usecase

import (
	"fmt"
	"strconv"
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type managerGame struct {
	storage   Storage
	presenter Presenter
	config    domain.ConfigGame
	handlers  map[domain.EventIncomingID]handlerEvent
}

type handlerEvent func(time time.Time, player *domain.Player, param string) error

func NewManagerGame(storage Storage, presenter Presenter, config domain.ConfigGame) *managerGame {
	m := &managerGame{
		storage:   storage,
		presenter: presenter,
		config:    config,
	}

	m.handlers = map[domain.EventIncomingID]handlerEvent{
		domain.EventIncomingRegisterPlayer: m.handlerRegisterPlayer,
		domain.EventIncomingEnterDungeon:   m.handleEnterDungeon,
		domain.EventIncomingKillMonster:    m.handleKillMonster,
		domain.EventIncomingNextFloor:      m.handleNextFloor,
		domain.EventIncomingPrevFloor:      m.handlePrevFloor,
		domain.EventIncomingEnterBossFloor: m.handleEnterBossFloor,
		domain.EventIncomingKillBoss:       m.handleKillBoss,
		domain.EventIncomingLeaveDungeon:   m.handleLeaveDungeon,
		domain.EventIncomingCannotContinue: m.handleCannotContinue,
		domain.EventIncomingReceiveDamage:  m.handleReceiveDamage,
		domain.EventIncomingRestoreHealth:  m.handleHealthRestored,
	}

	return m
}

func (m *managerGame) ProcessEvent(
	time time.Time, idPlayer int, idEvent domain.EventIncomingID, param string,
) error {
	if time.After(m.config.TimeClosed) {
		m.presenter.ShowMadeImpossible(time, idPlayer, int(idEvent))
		return fmt.Errorf("impossible move %d", idEvent)
	}

	player := m.storage.Get(idPlayer)
	if player == nil {
		player := domain.NewPlayer(idPlayer, m.config.CountFloors)
		m.storage.Save(player)

		if idEvent == domain.EventIncomingRegisterPlayer {
			m.presenter.ShowRegistered(time, idPlayer)
			return nil
		} else {
			player.State = domain.StatePlayerDisqual
			m.presenter.ShowDisqualified(time, idPlayer)
			return fmt.Errorf("player not registered")
		}
	}

	if player.State == domain.StatePlayerDisqual {
		return fmt.Errorf("player disqual")
	}

	if player.TimeEnterDungeon.IsZero() &&
		idEvent != domain.EventIncomingEnterDungeon {
		m.presenter.ShowMadeImpossible(time, player.Id, int(idEvent))
		return fmt.Errorf("player not enter")
	}

	if handler, ok := m.handlers[idEvent]; ok {
		return handler(time, player, param)
	}

	return fmt.Errorf("handler not mapped")
}

func (m *managerGame) handlerRegisterPlayer(
	time time.Time, player *domain.Player, param string,
) error {
	player.State = domain.StatePlayerDisqual
	m.presenter.ShowDisqualified(time, player.Id)
	return fmt.Errorf("player not registered")
}

func (m *managerGame) handleEnterDungeon(
	time time.Time, player *domain.Player, param string,
) error {
	if !player.TimeEnterDungeon.IsZero() || time.Before(m.config.TimeOpened) {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingEnterDungeon),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingEnterDungeon)
	}

	player.TimeEnterDungeon = time
	m.presenter.ShowEnteredDungeon(time, player.Id)
	return nil
}

func (m *managerGame) handleLeaveDungeon(
	time time.Time, player *domain.Player, param string,
) error {
	if player.TimeEnterDungeon.IsZero() {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingLeaveDungeon),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingLeaveDungeon)
	}

	player.TimeLeaveDungeon = time
	m.presenter.ShowLeftDungeon(time, player.Id)
	return nil
}

func (m *managerGame) handleNextFloor(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorCurrent == m.config.CountFloors {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingNextFloor),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingNextFloor)
	}

	player.FloorCurrent++
	timeEnter := player.FloorsMonsters[player.FloorCurrent-1].Floor.TimeEnter
	if timeEnter.IsZero() {
		timeEnter = time
	}

	m.presenter.ShowWentToFloorNext(time, player.Id)
	return nil
}

func (m *managerGame) handlePrevFloor(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorCurrent == domain.IndexFloorMin {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingPrevFloor),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingPrevFloor)
	}

	player.FloorCurrent--
	m.presenter.ShowWentToFloorPrev(time, player.Id)
	return nil
}

func (m *managerGame) handleEnterBossFloor(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorCurrent != m.config.CountFloors {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingEnterBossFloor),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingEnterBossFloor)
	}

	player.FloorBoss.TimeEnter = time
	m.presenter.ShowEnteredFloorBoss(time, player.Id)
	return nil
}

func (m *managerGame) handleKillMonster(
	time time.Time, player *domain.Player, param string,
) error {
	countKilled := player.FloorsMonsters[player.FloorCurrent-1].MonstersKilled
	if countKilled == m.config.CountMonstersPerFloors || player.FloorCurrent == m.config.CountFloors {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingKillMonster),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingKillMonster)
	}

	countKilled++
	if countKilled == m.config.CountMonstersPerFloors {
		player.FloorsMonsters[player.FloorCurrent-1].Floor.TimeClear = time
		player.FloorsMonsters[player.FloorCurrent-1].Floor.Cleared = true
	}

	m.presenter.ShowKilledMonster(time, player.Id)
	return nil
}

func (m *managerGame) handleKillBoss(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorBoss.Cleared || player.FloorCurrent != m.config.CountFloors {
		m.presenter.ShowMadeImpossible(
			time, player.Id, int(domain.EventIncomingKillBoss),
		)
		return fmt.Errorf("impossible move %d", domain.EventIncomingKillBoss)
	}

	player.FloorBoss.Cleared = true
	player.FloorBoss.TimeClear = time
	m.presenter.ShowKilledBoss(time, player.Id)
	return nil
}

func (m *managerGame) handleHealthRestored(
	time time.Time, player *domain.Player, param string,
) error {
	health, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	player.RestoreHealth(health)
	m.presenter.ShowRestoredHealth(time, player.Id, health)
	return nil
}

func (m *managerGame) handleReceiveDamage(
	time time.Time, player *domain.Player, param string,
) error {
	damage, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	event := player.ReceiveDamage(damage)
	m.presenter.ShowReceivedDamage(time, player.Id, damage)

	for _, idOut := range event {
		if domain.EventOutgoingID(idOut) == domain.EventOutgoingDead {
			m.presenter.ShowDead(time, player.Id)
		}
	}

	return nil
}

func (m *managerGame) handleCannotContinue(
	time time.Time, player *domain.Player, param string,
) error {
	return nil
}
