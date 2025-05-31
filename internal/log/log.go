package log

import (
	"fmt"
	"log/slog"
)

type Logger struct {
	*slog.Logger
}

func New() Logger {
	return Logger{
		Logger: slog.Default(),
	}
}

func (l *Logger) Infof(message string, params ...any) {
	l.Info(fmt.Sprintf(message, params...))
}

func (l *Logger) Errorf(message string, params ...any) {
	l.Error(fmt.Sprintf(message, params...))
}

func (l *Logger) Warnf(message string, params ...any) {
	l.Warn(fmt.Sprintf(message, params...))
}

func (l *Logger) Debugf(message string, params ...any) {
	l.Debug(fmt.Sprintf(message, params...))
}
