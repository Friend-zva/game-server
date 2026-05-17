package usecase

import (
	"time"
)

type Presenter interface {
	ShowWIP(time time.Time)
	ShowPlayerRegistered(time time.Time, idPlayer int)
}
