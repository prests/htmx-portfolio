package logging

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	options := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	loggerHandler := slog.NewTextHandler(os.Stdout, options)
	return slog.New(loggerHandler)
}
