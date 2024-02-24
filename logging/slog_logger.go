package logging

import (
	"context"
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
		jsonHandler := slog.NewJSONHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("env", string(env))})
		levelHandler := NewLevelHandler(slog.LevelInfo, jsonHandler) // production logs only info and above
		logger := slog.New(levelHandler)
		return &SLogger{
			logger: logger,
		}
	} else if env == Development {
		jsonHandler := slog.NewJSONHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("env", string(env))})
		levelHandler := NewLevelHandler(slog.LevelDebug, jsonHandler) // development logs everything
		logger := slog.New(levelHandler)
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
