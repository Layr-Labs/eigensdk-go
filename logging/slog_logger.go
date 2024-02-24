package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type SlogHandlerType string

const (
	SLogHandlerTypeJson     SlogHandlerType = "json" // prints debug and above
	SLogHandlerTypeJsonText SlogHandlerType = "text" // prints info and above
)

type SLogger struct {
	logger *slog.Logger
}

var _ Logger = (*SLogger)(nil)

func NewSlogTextLogger(env LogLevel) Logger {
	var handler slog.Handler
	if env == Production {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo, // prints info and above (default behavior but we add to be explicit)
		}).
			WithAttrs([]slog.Attr{slog.String("env", string(env))}) // add a default tag with env name
	} else if env == Development {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,            // we add source information in development mode only
			Level:     slog.LevelDebug, // prints debug and above
		}).
			WithAttrs([]slog.Attr{slog.String("env", string(env))}) // add a default tag with env name
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}
	logger := slog.New(handler)
	return &SLogger{
		logger: logger,
	}
}

func NewSlogJsonLogger(env LogLevel) Logger {
	var handler slog.Handler
	if env == Production {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo, // prints info and above (default behavior but we add to be explicit)
		}).
			WithAttrs([]slog.Attr{slog.String("env", string(env))}) // add a default tag with env name
	} else if env == Development {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,            // we add source information in development mode only
			Level:     slog.LevelDebug, // prints debug and above
		}).
			WithAttrs([]slog.Attr{slog.String("env", string(env))}) // add a default tag with env name
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}
	logger := slog.New(handler)
	return &SLogger{
		logger: logger,
	}
}

func (s SLogger) Debug(msg string, tags ...any) {
	s.logCorrectSource(slog.LevelDebug, msg, tags...)
}

func (s SLogger) Info(msg string, tags ...any) {
	s.logCorrectSource(slog.LevelInfo, msg, tags...)
}

func (s SLogger) Warn(msg string, tags ...any) {
	s.logCorrectSource(slog.LevelWarn, msg, tags...)
}

func (s SLogger) Error(msg string, tags ...any) {
	s.logCorrectSource(slog.LevelError, msg, tags...)
}

func (s SLogger) Fatal(msg string, tags ...any) {
	s.logCorrectSource(slog.LevelError, msg, tags...)
	os.Exit(1)
}

// logCorrectSource logs the message with the correct source information
// adapted from https://pkg.go.dev/log/slog#example-package-Wrapping
func (s SLogger) logCorrectSource(level slog.Level, msg string, tags ...any) {
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, logCorrectSource, Info/Debug/Warn/Error/Fatal]
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(tags...)
	_ = s.logger.Handler().Handle(context.Background(), r)
}

func (s SLogger) Debugf(template string, args ...interface{}) {
	s.logfCorrectSource(slog.LevelDebug, template, args...)
}

func (s SLogger) Infof(template string, args ...interface{}) {
	s.logfCorrectSource(slog.LevelInfo, template, args...)
}

func (s SLogger) Warnf(template string, args ...interface{}) {
	s.logfCorrectSource(slog.LevelWarn, template, args...)
}

func (s SLogger) Errorf(template string, args ...interface{}) {
	s.logfCorrectSource(slog.LevelError, template, args...)
}

func (s SLogger) Fatalf(template string, args ...interface{}) {
	s.logfCorrectSource(slog.LevelError, template, args...)
	os.Exit(1)
}

// logfCorrectSource logs the message with the correct source information
// adapted from https://pkg.go.dev/log/slog#example-package-Wrapping
func (s SLogger) logfCorrectSource(level slog.Level, template string, args ...interface{}) {
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, logfCorrectSource, Info/Debug/Warn/Error/Fatal]
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(template, args...), pcs[0])
	_ = s.logger.Handler().Handle(context.Background(), r)
}
