package usecase

import (
	"fmt"
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type gameManager struct {
	storage   Storage
	presenter Presenter
}

func NewGameManager(storage Storage, presenter Presenter) *gameManager {
	return &gameManager{
		storage:   storage,
		presenter: presenter,
	}
}

func (u *gameManager) ProcessEvent(
	time time.Time, idPlayer int, idEvent domain.EventIncomingID, param string,
) error {
	player := u.storage.Get(idPlayer)
	if player == nil {
		if domain.EventIncomingID(idEvent) == domain.EventIncomingRegisterPlayer {
			player := domain.NewPlayer(idPlayer, time)
			u.storage.Save(player)
			u.presenter.ShowPlayerRegistered(time, idPlayer)
		} else {
			return fmt.Errorf("player not registered")
		}
	}

	switch idEvent {
	default:
		u.presenter.ShowWIP(time)
	}

	return nil
}
