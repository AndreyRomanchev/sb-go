package logger

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	timeFormat string
	debug      bool
}

func New(timeformat string, debug bool) *Logger {
	return &Logger{
		timeFormat: timeformat,
		debug:      debug,
	}
}

func (l *Logger) Log(level string, s ...interface{}) {
	level = strings.ToLower(level)
	switch level {
	case "info", "warning":
		if l.debug {
			l.write(level, s...)
		}
	default:
		l.write(level, s...)
	}
}

func (l *Logger) write(level string, s ...interface{}) {
	fmt.Printf("[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
}
