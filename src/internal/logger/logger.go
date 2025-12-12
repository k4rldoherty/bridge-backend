// Package logger - a package for setting up the logger of the application
package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger() *Logger {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(l)
	return &Logger{l}
}
