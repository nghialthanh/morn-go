package logger

import (
	"fmt"
	"log/slog"
)

type ILogger interface {
	Info(msg string, args ...interface{})
	Infof(format string, args ...interface{})
	Error(msg string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(msg string, args ...interface{})
	Warnf(format string, args ...interface{})
	Debug(msg string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type SlogLogger struct{}

func NewFmtLogger() ILogger {
	return &SlogLogger{}
}

func (l *SlogLogger) Info(msg string, args ...interface{}) {
	slog.Info(msg, args...)
}

func (l *SlogLogger) Infof(format string, args ...interface{}) {
	slog.Info(fmt.Sprintf(format, args...))
}

func (l *SlogLogger) Error(msg string, args ...interface{}) {
	slog.Error(msg, args...)
}

func (l *SlogLogger) Errorf(format string, args ...interface{}) {
	slog.Error(fmt.Sprintf(format, args...))
}

func (l *SlogLogger) Warn(msg string, args ...interface{}) {
	slog.Warn(msg, args...)
}

func (l *SlogLogger) Warnf(format string, args ...interface{}) {
	slog.Warn(fmt.Sprintf(format, args...))
}

func (l *SlogLogger) Debug(msg string, args ...interface{}) {
	slog.Debug(msg, args...)
}

func (l *SlogLogger) Debugf(format string, args ...interface{}) {
	slog.Debug(fmt.Sprintf(format, args...))
}
