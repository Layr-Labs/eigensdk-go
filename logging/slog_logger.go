package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type SLogger struct {
	*slog.Logger
}

var _ Logger = (*SLogger)(nil)

// NewSlogTextLogger creates a new SLogger with a text handler
//
//		outputWriter is the writer to write the logs to (typically os.Stdout,
//		  but can also use a io.MultiWriter(os.Stdout, file) to write to multiple outputs)
//	 	handlerOptions is the options for the handler, such as
//		Level is the minimum level to log
//		AddSource if true, adds source information to the log
func NewSlogTextLogger(outputWriter io.Writer, handlerOpts *slog.HandlerOptions) *SLogger {
	handler := slog.NewTextHandler(outputWriter, handlerOpts)
	logger := slog.New(handler)
	return &SLogger{
		logger,
	}
}

// NewSlogJsonLogger creates a new SLogger with a Json handler
//
//		outputWriter is the writer to write the logs to (typically os.Stdout,
//		  but can also use a io.MultiWriter(os.Stdout, file) to write to multiple outputs)
//	 	handlerOptions is the options for the handler, such as
//		Level is the minimum level to log
//		AddSource if true, adds source information to the log
func NewSlogJsonLogger(outputWriter io.Writer, handlerOpts *slog.HandlerOptions) *SLogger {
	handler := slog.NewJSONHandler(outputWriter, handlerOpts)
	logger := slog.New(handler)
	return &SLogger{
		logger,
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
	_ = s.Handler().Handle(context.Background(), r)
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
	_ = s.Handler().Handle(context.Background(), r)
}
