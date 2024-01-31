package logging

import (
	"fmt"
	"log/slog"
	"os"
)

type SLogger struct {
	logger *slog.Logger
}

var _ Logger = (*SLogger)(nil)

func NewSlogLogger(env LogLevel) Logger {
	if env == Production {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		return &SLogger{
			logger: logger,
		}
	} else if env == Development {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		return &SLogger{
			logger: logger,
		}
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}
}

func (s SLogger) Debug(msg string, tags ...any) {
	s.logger.Debug(msg, tags...)
}

func (s SLogger) Info(msg string, tags ...any) {
	s.logger.Info(msg, tags...)
}

func (s SLogger) Warn(msg string, tags ...any) {
	s.logger.Warn(msg, tags...)
}

func (s SLogger) Error(msg string, tags ...any) {
	s.logger.Error(msg, tags...)
}

func (s SLogger) Fatal(msg string, tags ...any) {
	s.logger.Error(msg, tags...)
	os.Exit(1)
}

func (s SLogger) Debugf(template string, args ...interface{}) {
	s.logger.Debug(fmt.Sprintf(template, args...))
}

func (s SLogger) Infof(template string, args ...interface{}) {
	s.logger.Info(fmt.Sprintf(template, args...))
}

func (s SLogger) Warnf(template string, args ...interface{}) {
	s.logger.Warn(fmt.Sprintf(template, args...))
}

func (s SLogger) Errorf(template string, args ...interface{}) {
	s.logger.Error(fmt.Sprintf(template, args...))
}

func (s SLogger) Fatalf(template string, args ...interface{}) {
	s.logger.Error(fmt.Sprintf(template, args...))
	os.Exit(1)
}
