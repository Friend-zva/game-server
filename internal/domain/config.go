package domain

import (
	"time"
)

type ConfigGame struct {
	CountFloors            int
	CountMonstersPerFloors int
	TimeOpenAt             time.Time
	HoursDuration          int
}
