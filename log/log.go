/*
Package log contains the singleton object and helper functions for facilitating logging within the syft library.
*/
package log

import (
	"github.com/anchore/go-logger"
	"github.com/anchore/go-logger/adapter/discard"
	"github.com/anchore/go-logger/adapter/logrus"
)

// log is the singleton used to facilitate logging internally
var log = discard.New()

func SetLogger(logger logger.Logger) {
	log = logger
}

func DefaultLogger(cfg Config) error {
	switch cfg.Verbosity {
	case 1:
		cfg.Level = InfoLevel
	case 2:
		cfg.Level = DebugLevel
	default:
		cfg.Level = TraceLevel
	}
	c := logrus.Config{
		EnableConsole: (cfg.File == "" || cfg.Verbosity > 0) && !cfg.Quiet,
		FileLocation:  cfg.File,
		Level:         logger.Level(cfg.Level),
	}
	l, err := logrus.New(c)
	if err != nil {
		return err
	}

	log = l
	return nil
}

// Errorf takes a formatted template string and template arguments for the error logging level.
func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

// Error logs the given arguments at the error logging level.
func Error(args ...any) {
	log.Error(args...)
}

// Warnf takes a formatted template string and template arguments for the warning logging level.
func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}

// Warn logs the given arguments at the warning logging level.
func Warn(args ...any) {
	log.Warn(args...)
}

// Infof takes a formatted template string and template arguments for the info logging level.
func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

// Info logs the given arguments at the info logging level.
func Info(args ...any) {
	log.Info(args...)
}

// Debugf takes a formatted template string and template arguments for the debug logging level.
func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

// Debug logs the given arguments at the debug logging level.
func Debug(args ...any) {
	log.Debug(args...)
}

// Tracef takes a formatted template string and template arguments for the trace logging level.
func Tracef(format string, args ...any) {
	log.Tracef(format, args...)
}

// Trace logs the given arguments at the trace logging level.
func Trace(args ...any) {
	log.Trace(args...)
}

// WithFields returns a message logger with multiple key-value fields.
func WithFields(fields ...any) logger.MessageLogger {
	return log.WithFields(fields...)
}

// Nested returns a new logger with hard coded key-value pairs
func Nested(fields ...any) logger.Logger {
	return log.Nested(fields...)
}
