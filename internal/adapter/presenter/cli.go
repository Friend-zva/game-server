package cli

import (
	"fmt"
	"io"
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type presenterCLI struct {
	formatTime string
	out        io.Writer
}

func NewPresenterCLI(formatTime string, out io.Writer) *presenterCLI {
	return &presenterCLI{
		formatTime: formatTime,
		out:        out,
	}
}

func (p *presenterCLI) ShowRegistered(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] registered\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowDisqualified(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] is disqualified\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowEnteredDungeon(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] entered the dungeon\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowLeftDungeon(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] left the dungeon\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowWentToFloorNext(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] went to the next floor\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowWentToFloorPrev(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] went to the previous floor\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowEnteredFloorBoss(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] entered the boss's floor\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowCannotContinue(time time.Time, idPlayer int, reason string) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out,
		"[%s] Player [%d] cannot continue due to %s\n", strTime, idPlayer, reason,
	)
}

func (p *presenterCLI) ShowMadeImpossible(time time.Time, idPlayer, idEvent int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out,
		"[%s] Player [%d] makes imposible move [%d]\n", strTime, idPlayer, idEvent,
	)
}

func (p *presenterCLI) ShowKilledMonster(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] killed the monster\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowKilledBoss(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] killed the boss\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowRestoredHealth(time time.Time, idPlayer int, amount int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out,
		"[%s] Player [%d] has restored [%d] of health\n", strTime, idPlayer, amount,
	)
}

func (p *presenterCLI) ShowReceivedDamage(time time.Time, idPlayer, amount int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out,
		"[%s] Player [%d] recieved [%d] of damage\n", strTime, idPlayer, amount,
	)
}

func (p *presenterCLI) ShowDead(time time.Time, idPlayer int) {
	strTime := time.Format(p.formatTime)
	fmt.Fprintf(p.out, "[%s] Player [%d] is dead\n", strTime, idPlayer)
}

func (p *presenterCLI) ShowPreReportPlayer() {
	fmt.Fprintf(p.out, "Final report:\n")
}

func (p *presenterCLI) ShowReportPlayer(
	state domain.StatePlayer, idPlayer int,
	timeTotal, timeAvgFloor, timeBoss time.Duration,
	health int,
) {
	fmt.Fprintf(p.out, "[%s] %d [%s, %s, %s] HP:%d\n", state, idPlayer,
		formatDuration(timeTotal), formatDuration(timeAvgFloor), formatDuration(timeBoss),
		health,
	)
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours()) % 60
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
