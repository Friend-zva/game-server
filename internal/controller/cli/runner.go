package cli

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	dto "github.com/Friend-zva/game-server/internal/controller/dto"
	domain "github.com/Friend-zva/game-server/internal/domain"
)

type ManagerGame interface {
	ProcessEvent(
		time time.Time, idPlayer int, idEvent domain.EventIncomingID, param string,
	) error
	GenerateReport() error
}

type RunnerGame struct {
	logger  *slog.Logger
	manager ManagerGame
	decoder *dto.DecoderEvent
}

func NewRunnerGame(logger *slog.Logger, manager ManagerGame, formatTime string) *RunnerGame {
	return &RunnerGame{
		logger:  logger,
		manager: manager,
		decoder: dto.NewDecoderEvent(formatTime),
	}
}

func (r *RunnerGame) Execute(pathEvents string) error {
	file, err := os.Open(pathEvents)
	if err != nil {
		r.logger.Error("failed to open events file", "error", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			r.logger.Error("failed to close events file", "error", err)
		}
	}()

	return r.scanFromReader(file)
}

func (r *RunnerGame) scanFromReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		event, err := r.decoder.Execute(line)
		if err != nil {
			r.logger.Error("failed to decode event", "error", err, "line", line)
			continue
		}

		err = r.manager.ProcessEvent(event.Time, event.IdPlayer, event.IdEvent, event.Param)
		if err != nil {
			r.logger.Debug("failed to exec event", "error", err)
		}
	}

	err := r.manager.GenerateReport()
	if err != nil {
		return err
	}

	return nil
}
