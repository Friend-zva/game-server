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

func (p *presenterCLI) ShowWIP(time time.Time) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Work in progress\n", timeStr)
}

func (p *presenterCLI) ShowRegistered(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] registered\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowDisqualified(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] is disqualified\n", timeStr, idPlayer)
}

func (p *presenterCLI) ShowDamageReceived(time time.Time, idPlayer, damage int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] recieved [%d] of damage\n",
		timeStr,
		idPlayer,
		damage,
	)
}

func (p *presenterCLI) ShowDead(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] is dead\n", timeStr, idPlayer)
}
