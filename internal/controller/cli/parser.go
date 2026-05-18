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

type ManagerGame interface {
	ProcessEvent(
		time time.Time, idPlayer int, idEvent domain.EventIncomingID, param string,
	) error
	GenerateReport()
}

type parser struct {
	logger      *slog.Logger
	managerGame ManagerGame
	formatTime  string
}

func NewParser(logger *slog.Logger, managerGame ManagerGame, formatTime string) *parser {
	return &parser{
		logger:      logger,
		managerGame: managerGame,
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

		tokens := strings.Fields(parts[1])
		if len(tokens) < 2 {
			p.logger.Error("failed to parse event args", "error", err)
			continue
		}

		idPlayer, err := strconv.Atoi(tokens[0])
		if err != nil {
			p.logger.Error("failed to parse player id", "error", err)
			continue
		}

		event, err := strconv.Atoi(tokens[1])
		if err != nil {
			p.logger.Error("failed to parse event", "error", err)
			continue
		}

		var param string
		if len(tokens) >= 3 {
			param = strings.Join(tokens[2:], " ")
		}

		idEvent := domain.EventIncomingID(event)
		err = p.managerGame.ProcessEvent(time, idPlayer, idEvent, param)
		if err != nil {
			p.logger.Debug("failed to exec event", "error", err)
			continue
		}
	}

	p.managerGame.GenerateReport()
	return nil
}
