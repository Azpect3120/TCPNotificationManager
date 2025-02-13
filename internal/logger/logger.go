package logger

import (
	"fmt"
	"time"
)

type LogLevel string

// The different levels of logging that can be used.
const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

// Function synmbol used to configure the logger.
type LoggerOptsFunc func(*LoggerOpts)

// Options used to configure the logger.
type LoggerOpts struct {
	// The level of logging to be used when none
	// is specified.
	DefaultLevel LogLevel

	// Whether or not to include a timestamp in the
	// log output.
	Timestamp bool
}

// Provide a default log level for the logger.
func WithDefaultLevel(level LogLevel) LoggerOptsFunc {
	return func(opts *LoggerOpts) {
		opts.DefaultLevel = level
	}
}

// Enable timestamps in the log output.
func WithTimestamp() LoggerOptsFunc {
	return func(opts *LoggerOpts) {
		opts.Timestamp = true
	}
}

// Defines the default logger options, if they are not
// provided by the user.
func defaultLoggerOpts() LoggerOpts {
	return LoggerOpts{
		DefaultLevel: INFO,
		Timestamp:    false,
	}
}

// Logger is a simple logging package that allows for
// different levels of logging to be used. Each log
// message will include the log level, and the message
// that was passed to the logger.
type Logger struct {
	// Logger options.
	Opts LoggerOpts
}

// Log a message at the specified level. If no level is
// provided, the default level will be used.
//
// TODO: If throughput is a concern, this logger can be
// optimized by using a queue to store log messages, and
// then have a separate goroutine that reads from the
// queue and writes to the output.
//
// TODO: I would also like to implement log files, so this
// logger can be used to write to multiple outputs using
// goroutine.
func NewLogger(opts ...LoggerOptsFunc) *Logger {
	logger := &Logger{
		Opts: defaultLoggerOpts(),
	}

	// Apply options to the logger.
	for _, opt := range opts {
		opt(&logger.Opts)
	}

	return logger
}

// Log a message at the specified level. If no level is
// provided, the default level will be used.
//
// The only way to allow for optional params is to use the
// ... operator. This will allow for the function to accept
// 0 or more LogLevel arguments. Only the first argument
// will be used, if it exists. If any more arguments are
// provided, they will be ignored.
func (l *Logger) Log(message string, level ...LogLevel) {
	var logLevel LogLevel
	if len(level) > 0 {
		logLevel = level[0]
	} else {
		logLevel = l.Opts.DefaultLevel
	}

	if l.Opts.Timestamp {
		fmt.Printf("[%s] [%s] %s", logLevel, time.Now().Format(time.RFC3339), message)
	} else {
		fmt.Printf("[%s] %s", logLevel, message)
	}
}
