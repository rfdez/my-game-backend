package logger

// Logger is an interface for logging.
type Logger interface {
	Debug(string)
	Info(string)
	Error(string)
	Fatal(string)
}

//go:generate mockery --case=snake --outpkg=loggermocks --output=loggermocks --name=Logger
