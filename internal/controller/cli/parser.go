package cli

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type GameManager interface {
	ProcessEvent(
		time time.Time, idPlayer int, idEvent domain.EventIncomingID, param string,
	) error
}

type parser struct {
	logger      *slog.Logger
	gameManager GameManager
	formatTime  string
}

func NewParser(logger *slog.Logger, gameManager GameManager, formatTime string) *parser {
	return &parser{
		logger:      logger,
		gameManager: gameManager,
		formatTime:  formatTime,
	}
}

func (p *parser) Run(pathEvents string) error {
	file, err := os.Open(pathEvents)
	if err != nil {
		p.logger.Error("failed to open events file", "error", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, "] ")
		timeRaw := strings.TrimPrefix(parts[0], "[")
		time, err := time.Parse(p.formatTime, timeRaw)
		if err != nil {
			p.logger.Error("failed to parse time", "error", err)
			continue
		}

		rest := strings.Fields(parts[1])
		if len(rest) < 2 {
			p.logger.Error("failed to parse event", "error", err)
			continue
		}

		idPlayer, err := strconv.Atoi(rest[0])
		if err != nil {
			p.logger.Error("failed to parse player id", "error", err)
			continue
		}

		event, err := strconv.Atoi(rest[1])
		if err != nil {
			p.logger.Error("failed to parse event", "error", err)
			continue
		}

		var param string
		if len(rest) == 3 {
			param = rest[2]
		}

		idEvent := domain.EventIncomingID(event)
		err = p.gameManager.ProcessEvent(time, idPlayer, idEvent, param)
		if err != nil {
			p.logger.Error("failed to exec event", "error", err)
			continue
		}
	}

	return nil
}
