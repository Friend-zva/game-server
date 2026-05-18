package usecase

import (
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type Presenter interface {
	ShowRegistered(time time.Time, idPlayer int)                    // 1
	ShowEnteredDungeon(time time.Time, idPlayer int)                // 2
	ShowKilledMonster(time time.Time, idPlayer int)                 // 3
	ShowWentToFloorNext(time time.Time, idPlayer int)               // 4
	ShowWentToFloorPrev(time time.Time, idPlayer int)               // 5
	ShowEnteredFloorBoss(time time.Time, idPlayer int)              // 6
	ShowKilledBoss(time time.Time, idPlayer int)                    // 7
	ShowLeftDungeon(time time.Time, idPlayer int)                   // 8
	ShowCannotContinue(time time.Time, idPlayer int, reason string) // 9
	ShowRestoredHealth(time time.Time, idPlayer, amount int)        // 10
	ShowReceivedDamage(time time.Time, idPlayer, amount int)        // 11
	ShowDisqualified(time time.Time, idPlayer int)                  // 31
	ShowDead(time time.Time, idPlayer int)                          // 32
	ShowMadeImpossible(time time.Time, idPlayer, idEvent int)       // 33
	ShowPreReportPlayer()
	ShowReportPlayer(state domain.StatePlayer, idPlayer int,
		timeTotal, timeAvgFloor, timeBoss time.Duration,
		health int)
}
