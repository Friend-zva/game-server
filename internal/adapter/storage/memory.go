package storage

import (
	domain "github.com/Friend-zva/game-server/internal/domain"
)

type storageMemory struct {
	players map[int]*domain.Player
}

func NewStorageMemory() *storageMemory {
	return &storageMemory{
		players: make(map[int]*domain.Player),
	}
}

func (s *storageMemory) Get(id int) *domain.Player {
	return s.players[id]
}

func (s *storageMemory) Save(player *domain.Player) {
	s.players[player.Id] = player
}
