package cli

import (
	"fmt"
	"time"
)

type cliPresenter struct {
	formatTime string
}

func NewCLIPresenter(formatTime string) *cliPresenter {
	return &cliPresenter{
		formatTime: formatTime,
	}
}

func (p *cliPresenter) ShowWIP(time time.Time) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Work in progress\n", timeStr)
}

func (p *cliPresenter) ShowPlayerRegistered(time time.Time, idPlayer int) {
	timeStr := time.Format(p.formatTime)
	fmt.Printf("[%s] Player [%d] registered\n", timeStr, idPlayer)
}
