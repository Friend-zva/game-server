package domain

import (
	"time"
)

const IndexFloorMin = 1
const HealthMax = 100

type ConfigGame struct {
	CountFloors            int
	CountMonstersPerFloors int
	TimeOpenAt             time.Time
	HoursDuration          int
}
