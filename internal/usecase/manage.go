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
		domain.EventIncomingReceiveDamage: m.handleReceiveDamage,
	}

	return m
}

func (m *managerGame) handleReceiveDamage(
	time time.Time, player *domain.Player, param string,
) error {
	damage, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	event := player.TakeDamage(damage)
	m.presenter.ShowDamageReceived(time, player.Id, damage)

	for _, idOut := range event {
		if domain.EventOutgoingID(idOut) == domain.EventOutgoingDead {
			m.presenter.ShowDead(time, player.Id)
		}
	}
	return nil
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

	m.presenter.ShowWIP(time)
	return nil
}
