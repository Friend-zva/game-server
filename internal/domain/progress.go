package domain

import (
	"time"
)

type Floor struct {
	Cleared   bool
	TimeEnter time.Time
	TimeClear time.Time
}

type FloorMonsters struct {
	Floor          Floor
	MonstersKilled int
}
