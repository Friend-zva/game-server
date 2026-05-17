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
		domain.EventIncomingEnterDungeon: m.handleEnterDungeon,
		// domain.EventIncomingKillMonster:    m.handleKillMonster,
		domain.EventIncomingNextFloor:      m.handleNextFloor,
		domain.EventIncomingPrevFloor:      m.handlePrevFloor,
		domain.EventIncomingEnterBossFloor: m.handleEnterBossFloor,
		// domain.EventIncomingKillBoss:       m.handleKillBoss,
		domain.EventIncomingLeaveDungeon: m.handleLeaveDungeon,
		// domain.EventIncomingCannotContinue: m.handleCannotContinue,
		domain.EventIncomingReceiveDamage: m.handleReceiveDamage,
		domain.EventIncomingRestoreHealth: m.handleHealthRestored,
	}

	return m
}

func (m *managerGame) ProcessEvent(
	time time.Time, idPlayer int, idEvent domain.EventIncomingID, param string,
) error {
	player := m.storage.Get(idPlayer)
	if player == nil {
		player := domain.NewPlayer(idPlayer, time)
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

	if handler, ok := m.handlers[idEvent]; ok {
		return handler(time, player, param)
	}

	return fmt.Errorf("handler not mapped")
}

func (m *managerGame) handleEnterDungeon(
	time time.Time, player *domain.Player, param string,
) error {
	player.EnteredDungeon = true
	m.presenter.ShowEnteredDungeon(time, player.Id)
	return nil
}

func (m *managerGame) handleLeaveDungeon(
	time time.Time, player *domain.Player, param string,
) error {
	player.EnteredDungeon = false
	m.presenter.ShowLeftDungeon(time, player.Id)
	return nil
}

func (m *managerGame) handleNextFloor(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorCurrent == m.config.CountFloors {
		m.presenter.ShowMadeImposible(
			time, player.Id, int(domain.EventIncomingNextFloor),
		)
		return fmt.Errorf("imposible move")
	}

	player.FloorCurrent++
	m.presenter.ShowWentToFloorNext(time, player.Id)
	return nil
}

func (m *managerGame) handlePrevFloor(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorCurrent == domain.IndexFloorMin {
		m.presenter.ShowMadeImposible(
			time, player.Id, int(domain.EventIncomingPrevFloor),
		)
		return fmt.Errorf("imposible move")
	}

	player.FloorCurrent--
	m.presenter.ShowWentToFloorPrev(time, player.Id)
	return nil
}

func (m *managerGame) handleEnterBossFloor(
	time time.Time, player *domain.Player, param string,
) error {
	if player.FloorCurrent != m.config.CountFloors-1 {
		m.presenter.ShowMadeImposible(
			time, player.Id, int(domain.EventIncomingEnterBossFloor),
		)
		return fmt.Errorf("imposible move")
	}

	player.FloorCurrent++
	m.presenter.ShowEnteredFloorBoss(time, player.Id)
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
