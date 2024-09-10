package logconfig

import (
	"log/slog"
	"os"
)

const (
	ErrorKey       = "err"
	ElapsedTimeKey = "elapsed_time"
	PathKey        = "path"
)

// SetDefaultLogger configures the default logger for qcc-study-hall-gopher applications.
func SetDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)
}
