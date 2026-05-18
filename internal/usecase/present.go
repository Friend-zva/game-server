package usecase

import (
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type Presenter interface {
	ShowRegistered(time time.Time, idPlayer int) error                    // 1
	ShowEnteredDungeon(time time.Time, idPlayer int) error                // 2
	ShowKilledMonster(time time.Time, idPlayer int) error                 // 3
	ShowWentToFloorNext(time time.Time, idPlayer int) error               // 4
	ShowWentToFloorPrev(time time.Time, idPlayer int) error               // 5
	ShowEnteredFloorBoss(time time.Time, idPlayer int) error              // 6
	ShowKilledBoss(time time.Time, idPlayer int) error                    // 7
	ShowLeftDungeon(time time.Time, idPlayer int) error                   // 8
	ShowCannotContinue(time time.Time, idPlayer int, reason string) error // 9
	ShowRestoredHealth(time time.Time, idPlayer, amount int) error        // 10
	ShowReceivedDamage(time time.Time, idPlayer, amount int) error        // 11
	ShowDisqualified(time time.Time, idPlayer int) error                  // 31
	ShowDead(time time.Time, idPlayer int) error                          // 32
	ShowMadeImpossible(                                                   // 33
		time time.Time, idPlayer int, idEvent domain.EventIncomingID) error
	ShowPreReportPlayer() error
	ShowReportPlayer(state domain.StatePlayer, idPlayer int,
		timeTotal, timeAvgFloor, timeBoss time.Duration,
		health int) error
}
