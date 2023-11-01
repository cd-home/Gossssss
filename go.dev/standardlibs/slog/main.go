package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("ok")
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	logger.Info("OK")
}
