package usecase

import (
	domain "github.com/Friend-zva/game-server/internal/domain"
)

type Storage interface {
	Get(id int) *domain.Player
	Save(player *domain.Player)
}
