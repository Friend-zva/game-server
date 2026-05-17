package main

import (
	"flag"
	"fmt"
	"os"

	presenter "github.com/Friend-zva/game-server/internal/adapter/presenter"
	storage "github.com/Friend-zva/game-server/internal/adapter/storage"
	clicontroller "github.com/Friend-zva/game-server/internal/controller/cli"
	usecase "github.com/Friend-zva/game-server/internal/usecase"
	logger "github.com/Friend-zva/game-server/platform/logger"
)

func run() error {
	var pathConfig, pathEvents string
	flag.StringVar(&pathConfig, "config", "config.json", "game configuration file")
	flag.StringVar(&pathEvents, "events", "events", "events file")
	flag.Parse()
	// cfg := config.MustLoad(pathConfig)

	logger := logger.MustMakeLogger("INFO")
	logger.Debug("debug messages are enabled")

	formatTime := "15:04:05"

	repo := storage.NewMemoryStorage()
	presenter := presenter.NewCLIPresenter(formatTime)
	gameManager := usecase.NewGameManager(repo, presenter)

	app := clicontroller.NewParser(logger, gameManager, formatTime)
	err := app.Run(pathEvents)
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
