package storage

import (
	domain "github.com/Friend-zva/game-server/internal/domain"
)

type memoryStorage struct {
	players map[int]*domain.Player
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		players: make(map[int]*domain.Player),
	}
}

func (r *memoryStorage) Get(id int) *domain.Player {
	return r.players[id]
}

func (r *memoryStorage) Save(player *domain.Player) {
	r.players[player.Id] = player
}
