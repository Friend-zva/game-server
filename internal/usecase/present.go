package usecase

import (
	"time"
)

type Presenter interface {
	ShowWIP(time time.Time)
	ShowRegistered(time time.Time, idPlayer int)
	ShowDisqualified(time time.Time, idPlayer int)
	ShowDamageReceived(time time.Time, playerID, damage int)
	ShowDead(time time.Time, playerID int)
}
