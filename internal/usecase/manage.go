package usecase

import (
	"fmt"
	"sort"
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
		m.presenter.ShowMadeImpossible(time, idPlayer, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, domain.ErrImpossibleMove)
	}

	player := m.storage.Get(idPlayer)
	if player == nil {
		player = domain.NewPlayer(idPlayer, m.config.CountFloors)
		m.storage.Save(player)

		if idEvent == domain.EventIncomingRegisterPlayer {
			m.presenter.ShowRegistered(time, idPlayer)
			return nil
		}

		player.State = domain.StatePlayerDisqual
		m.presenter.ShowDisqualified(time, idPlayer)
		m.storage.Save(player)
		return fmt.Errorf("impossible move %d: %w", idEvent, domain.ErrPlayerNotRegister)
	}

	if player.State == domain.StatePlayerDisqual {
		return domain.ErrPlayerDisqual
	}

	if player.TimeEnterDungeon.IsZero() &&
		idEvent != domain.EventIncomingEnterDungeon {
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, domain.ErrPlayerNotEnter)
	}

	if handler, ok := m.handlers[idEvent]; ok {
		err := handler(time, player, param)
		m.storage.Save(player)
		return err
	}

	return fmt.Errorf("handler not mapped %d", idEvent)
}

func (m *managerGame) GenerateReport() {
	players := m.storage.GetAll()

	sort.Slice(players, func(i, j int) bool {
		return players[i].Id < players[j].Id
	})

	m.presenter.ShowPreReportPlayer()

	for _, p := range players {
		timeTotal := p.TimeLeaveDungeon.Sub(p.TimeEnterDungeon)

		timeAvgFloor := time.Duration(0)
		countFloorsMonsters := len(p.FloorsMonsters)
		countCleared := 0

		for _, f := range p.FloorsMonsters {
			if f.Floor.Cleared {
				timeAvgFloor += f.Floor.TimeClear.Sub(f.Floor.TimeEnter)
				countCleared++
			}
		}

		if countFloorsMonsters > 0 && countCleared > 0 {
			timeAvgFloor /= time.Duration(countFloorsMonsters)
		}

		timeBoss := p.FloorBoss.TimeClear.Sub(p.FloorBoss.TimeEnter)

		m.presenter.ShowReportPlayer(
			p.State, p.Id, timeTotal, timeAvgFloor, timeBoss, p.Health,
		)
	}
}

func (m *managerGame) handlerRegisterPlayer(
	time time.Time, player *domain.Player, param string,
) error {
	idEvent := domain.EventIncomingRegisterPlayer
	m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
	return domain.ErrImpossibleMove
}

func (m *managerGame) handleEnterDungeon(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.EnterDungeon(time, m.config.TimeOpened)
	if err != nil {
		idEvent := domain.EventIncomingEnterDungeon
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

	m.presenter.ShowEnteredDungeon(time, player.Id)
	return nil
}

func (m *managerGame) handleLeaveDungeon(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.LeaveDungeon(time)
	if err != nil {
		idEvent := domain.EventIncomingLeaveDungeon
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

	m.presenter.ShowLeftDungeon(time, player.Id)
	return nil
}

func (m *managerGame) handleNextFloor(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.NextFloor(time, m.config.CountFloors)
	if err != nil {
		idEvent := domain.EventIncomingNextFloor
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

	m.presenter.ShowWentToFloorNext(time, player.Id)
	return nil
}

func (m *managerGame) handlePrevFloor(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.PrevFloor()
	if err != nil {
		idEvent := domain.EventIncomingPrevFloor
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

	m.presenter.ShowWentToFloorPrev(time, player.Id)
	return nil
}

func (m *managerGame) handleEnterBossFloor(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.EnterBossFloor(time, m.config.CountFloors)
	if err != nil {
		idEvent := domain.EventIncomingEnterBossFloor
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

	m.presenter.ShowEnteredFloorBoss(time, player.Id)
	return nil
}

func (m *managerGame) handleKillMonster(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.KillMonster(time, m.config.CountMonstersPerFloors, m.config.CountFloors)
	if err != nil {
		idEvent := domain.EventIncomingKillMonster
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

	m.presenter.ShowKilledMonster(time, player.Id)
	return nil
}

func (m *managerGame) handleKillBoss(
	time time.Time, player *domain.Player, param string,
) error {
	err := player.KillBoss(time, m.config.CountFloors)
	if err != nil {
		idEvent := domain.EventIncomingKillBoss
		m.presenter.ShowMadeImpossible(time, player.Id, idEvent)
		return fmt.Errorf("impossible move %d: %w", idEvent, err)
	}

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

	player.ReceiveDamage(damage, time)
	m.presenter.ShowReceivedDamage(time, player.Id, damage)

	if player.State == domain.StatePlayerFail {
		m.presenter.ShowDead(time, player.Id)
	}

	return nil
}

func (m *managerGame) handleCannotContinue(
	time time.Time, player *domain.Player, param string,
) error {
	player.CannotContinue(time)
	m.presenter.ShowCannotContinue(time, player.Id, param)
	return nil
}
