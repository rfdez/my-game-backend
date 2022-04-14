package zerolog

import (
	"os"

	zero "github.com/rs/zerolog"
)

// Logger implements the Logger interface.
type Logger struct {
	logger zero.Logger
}

// NewLogger initializes a new instance of Logger.
func NewLogger() Logger {
	return Logger{
		logger: zero.New(os.Stdout),
	}
}

// Logger returns the logger instance.
func (l Logger) Logger() zero.Logger {
	return l.logger
}

// Debug implements the Logger interface.
func (l Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Info implements the Logger interface.
func (l Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Error implements the Logger interface.
func (l Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

// Fatal implements the Logger interface.
func (l Logger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}
