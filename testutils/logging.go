package testutils

import (
	"log/slog"
	"os"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

// NewTestLogger is just a utility function to create a logger for testing
// It returns an slog textHandler logger that outputs to os.Stdout, with source code information and debug level
func NewTestLogger() logging.Logger {
	// we don't use colors because the test output panel in vscode doesn't support them
	return logging.NewTextSLogger(
		os.Stdout,
		&logging.SLoggerOptions{AddSource: true, Level: slog.LevelDebug, NoColor: true},
	)
}
