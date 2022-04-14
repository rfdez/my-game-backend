package logger

// Logger is an interface for logging.
type Logger interface {
	Debug(string)
	Info(string)
	Error(string)
	Fatal(string)
}
