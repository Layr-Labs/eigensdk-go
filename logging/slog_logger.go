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
	var levelHandler *LevelHandler
	if env == Production {
		textHandler := slog.NewTextHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("env", string(env))})
		levelHandler = NewLevelHandler(slog.LevelInfo, textHandler) // production logs only info and above
	} else if env == Development {
		// we add source information in development mode only
		textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}).WithAttrs([]slog.Attr{slog.String("env", string(env))})
		levelHandler = NewLevelHandler(slog.LevelDebug, textHandler) // development logs everything
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}
	logger := slog.New(levelHandler)
	return &SLogger{
		logger: logger,
	}
}

func NewSlogJsonLogger(env LogLevel) Logger {
	var levelHandler *LevelHandler
	if env == Production {
		jsonHandler := slog.NewJSONHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("env", string(env))})
		levelHandler = NewLevelHandler(slog.LevelInfo, jsonHandler) // production logs only info and above
	} else if env == Development {
		// we add source information in development mode only
		jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}).WithAttrs([]slog.Attr{slog.String("env", string(env))})
		levelHandler = NewLevelHandler(slog.LevelDebug, jsonHandler) // development logs everything
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}
	logger := slog.New(levelHandler)
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

// Taken from https://pkg.go.dev/log/slog@master#example-Handler-LevelHandler
// A LevelHandler wraps a Handler with an Enabled method
// that returns false for levels below a minimum.
// This permits our logger to only log messages of a certain level or higher
type LevelHandler struct {
	level   slog.Leveler
	handler slog.Handler
}

// NewLevelHandler returns a LevelHandler with the given level.
// All methods except Enabled delegate to h.
func NewLevelHandler(level slog.Leveler, h slog.Handler) *LevelHandler {
	// Optimization: avoid chains of LevelHandlers.
	if lh, ok := h.(*LevelHandler); ok {
		h = lh.Handler()
	}
	return &LevelHandler{level, h}
}

// Enabled implements Handler.Enabled by reporting whether
// level is at least as large as h's level.
func (h *LevelHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

// Handle implements Handler.Handle.
func (h *LevelHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

// WithAttrs implements Handler.WithAttrs.
func (h *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithAttrs(attrs))
}

// WithGroup implements Handler.WithGroup.
func (h *LevelHandler) WithGroup(name string) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithGroup(name))
}

// Handler returns the Handler wrapped by h.
func (h *LevelHandler) Handler() slog.Handler {
	return h.handler
}
