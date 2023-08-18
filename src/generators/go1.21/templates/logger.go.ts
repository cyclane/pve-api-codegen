export default function () {
  return `package pve

import (
	"fmt"
	"io"
	"os"
)

// LeveledLogger is the logger interface
// required for a PVE API Client.
type LeveledLogger interface {
	Debugf(format string, v ...any)
	Infof(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
}

type LoggerLevel uint

const (
	LevelNone LoggerLevel = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// DefaultLogger is a basic LeveledLogger that is suitable in most use-cases.
type DefaultLogger struct {
	Level     LoggerLevel
	outWriter io.Writer
	errWriter io.Writer
}

// UseOutWriter sets the outWriter of the DefaultLogger.
func (dl *DefaultLogger) UseOutWriter(outWriter io.Writer) *DefaultLogger {
	dl.outWriter = outWriter
	return dl
}

// UseErrWriter sets the errWriter of the DefaultLogger.
func (dl *DefaultLogger) UseErrWriter(errWriter io.Writer) *DefaultLogger {
	dl.errWriter = errWriter
	return dl
}

func (dl *DefaultLogger) Debugf(format string, v ...any) {
	if dl.Level >= LevelDebug {
		fmt.Fprintf(dl.outWriter, "[DEBUG] "+format+"\\n", v...)
	}
}

func (dl *DefaultLogger) Infof(format string, v ...any) {
	if dl.Level >= LevelDebug {
		fmt.Fprintf(dl.outWriter, "[INFO]  "+format+"\\n", v...)
	}
}

func (dl *DefaultLogger) Warnf(format string, v ...any) {
	if dl.Level >= LevelDebug {
		fmt.Fprintf(dl.errWriter, "[WARN]  "+format+"\\n", v...)
	}
}

func (dl *DefaultLogger) Errorf(format string, v ...any) {
	if dl.Level >= LevelDebug {
		fmt.Fprintf(dl.errWriter, "[ERROR] "+format+"\\n", v...)
	}
}

// NewDefaultLogger creates a new DefaultLogger with the given LoggerLevel.
func NewDefaultLogger(level LoggerLevel) *DefaultLogger {
	return &DefaultLogger{
		Level:     level,
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
}
`;
}
