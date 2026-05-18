package cli

import (
	"fmt"
	"time"
)

type presenterCLI struct {
	formatTime string
}

func NewPresenterCLI(formatTime string) *presenterCLI {
	return &presenterCLI{
		formatTime: formatTime,
	}
}

func (p *presenterCLI) ShowRegistered(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] registered\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowDisqualified(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] is disqualified\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowEnteredDungeon(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] entered the dungeon\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowLeftDungeon(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] left the dungeon\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowWentToFloorNext(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] went to the next floor\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowWentToFloorPrev(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] went to the previous floor\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowEnteredFloorBoss(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] entered the boss's floor\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowCannotContinue(time time.Time, idPlayer int, reason string) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf(
		"[%s] Player [%d] cannot continue due to %s\n", timeStr, idPlayer, reason,
	)
}

func (p *presenterCLI) ShowMadeImpossible(time time.Time, idPlayer, idEvent int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf(
		"[%s] Player [%d] makes imposible move [%d]\n", timeStr, idPlayer, idEvent,
	)
}

func (p *presenterCLI) ShowKilledMonster(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] killed the monster\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowKilledBoss(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] killed the boss\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowRestoredHealth(time time.Time, idPlayer int, amount int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf(
		"[%s] Player [%d] has restored [%d] of health\n", timeStr, idPlayer, amount,
	)
}

func (p *presenterCLI) ShowReceivedDamage(time time.Time, idPlayer, amount int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf(
		"[%s] Player [%d] recieved [%d] of damage\n", timeStr, idPlayer, amount,
	)
}

func (p *presenterCLI) ShowDead(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] is dead\n", timeStr, idPlayer)
}
