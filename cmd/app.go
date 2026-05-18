package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	config "github.com/Friend-zva/game-server/config"
	presenter "github.com/Friend-zva/game-server/internal/adapter/presenter"
	storage "github.com/Friend-zva/game-server/internal/adapter/storage"
	clicontroller "github.com/Friend-zva/game-server/internal/controller/cli"
	domain "github.com/Friend-zva/game-server/internal/domain"
	usecase "github.com/Friend-zva/game-server/internal/usecase"
	logger "github.com/Friend-zva/game-server/platform/logger"
)

func run() error {
	var pathConfig, pathEvents string
	flag.StringVar(&pathConfig, "config", "config/config.json", "game configuration file")
	flag.StringVar(&pathEvents, "events", "events", "events file")
	flag.Parse()
	cfg := config.MustLoad(pathConfig)

	logger := logger.MustMakeLogger("DEBUG")
	logger.Debug("debug messages are enabled")

	formatTime := "15:04:05"

	timeOpenAt, err := time.Parse(formatTime, cfg.OpenAt)
	if err != nil {
		return fmt.Errorf("cannot parse game config: %w", err)
	}

	configGame := domain.ConfigGame{
		CountFloors:            cfg.Floors,
		CountMonstersPerFloors: cfg.Monsters,
		TimeOpened:             timeOpenAt,
		TimeClosed:             timeOpenAt.Add(time.Duration(cfg.Duration) * time.Hour),
		HoursDuration:          cfg.Duration,
	}

	storage := storage.NewStorageMemory()
	presenter := presenter.NewPresenterCLI(formatTime)
	managerGame := usecase.NewManagerGame(storage, presenter, configGame)

	app := clicontroller.NewParser(logger, managerGame, formatTime)
	err = app.Run(pathEvents)
	if err != nil {
		return fmt.Errorf("cannot run game: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
