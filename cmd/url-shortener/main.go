package main

import (
	"fmt"
	"log/slog"
	"os"
	"urlShortener/internal/config"
	"urlShortener/internal/storage"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv **
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("starting app", slog.String("env", cfg.Env))
	log.Debug("debug messages are turned on")
	// todo: init storage: sqlite3
	storage, err := storage.sqlite.New(cfg.StoragePath)
	// TODO: init router: chi, chi-render
	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
