package zerologlogger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger implements the Logger interface.
type logger struct {
	logger zerolog.Logger
}

// NewLogger initializes a new instance of Logger.
func NewLogger() logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	zerologger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	return logger{
		logger: zerologger,
	}
}

// Debug implements the Logger interface.
func (l logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Info implements the Logger interface.
func (l logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Error implements the Logger interface.
func (l logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

// Fatal implements the Logger interface.
func (l logger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}
