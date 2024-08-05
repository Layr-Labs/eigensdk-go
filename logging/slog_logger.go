package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/lmittmann/tint"
)

type SLogger struct {
	*slog.Logger
}

var _ Logger = (*SLogger)(nil)

// SLoggerOptions are options when creating a new SLogger.
// A zero Options consists entirely of default values.
//
// SLoggerOptions are an extension of [slog.HandlerOptions].
type SLoggerOptions struct {
	// Enable source code location (Default: false)
	AddSource bool

	// Minimum level to log (Default: slog.LevelInfo)
	Level slog.Leveler

	// ReplaceAttr is called to rewrite each non-group attribute before it is logged.
	// See https://pkg.go.dev/log/slog#HandlerOptions for details.
	ReplaceAttr func(groups []string, attr slog.Attr) slog.Attr

	// Time format (Default: time.StampMilli)
	// only supported with text handler
	TimeFormat string

	// Disable color (Default: false)
	// only supported with text handler (no color in json)
	NoColor bool
}

// default SLogger options are used when no options are provided
// they are the development options (debug logs with source)
var defaultTintOptions = tint.Options{
	AddSource:   false,
	Level:       slog.LevelInfo,
	ReplaceAttr: nil,
	TimeFormat:  time.StampMilli,
	NoColor:     false,
}

// NewSlogTextLogger creates a new SLogger with a text handler
// Default behavior is colored log outputs. To disable colors, set opts.NoColor to true.
func NewTextSLogger(outputWriter io.Writer, opts *SLoggerOptions) *SLogger {
	var tintOptions tint.Options
	if opts == nil {
		tintOptions = defaultTintOptions
	} else {
		tintOptions = tint.Options{
			AddSource:   opts.AddSource,
			Level:       opts.Level,
			ReplaceAttr: opts.ReplaceAttr,
			TimeFormat:  opts.TimeFormat,
			NoColor:     opts.NoColor,
		}
	}
	handler := tint.NewHandler(outputWriter, &tintOptions)
	logger := slog.New(handler)
	return &SLogger{
		logger,
	}
}

// NewSlogJsonLogger creates a new SLogger with a Json handler
// Currently colors are not supported with json handler. If colors are required,
// use NewTextSLogger instead.
func NewJsonSLogger(outputWriter io.Writer, opts *SLoggerOptions) *SLogger {
	var handlerOpts *slog.HandlerOptions
	if opts == nil {
		handlerOpts = &slog.HandlerOptions{}
	} else {
		handlerOpts = &slog.HandlerOptions{
			AddSource:   opts.AddSource,
			Level:       opts.Level,
			ReplaceAttr: opts.ReplaceAttr,
		}
	}
	handler := slog.NewJSONHandler(outputWriter, handlerOpts)
	logger := slog.New(handler)
	return &SLogger{
		logger,
	}
}

// NewSlogTextLogger creates a new SLogger with a text handler
//
//		outputWriter is the writer to write the logs to (typically os.Stdout,
//		  but can also use a io.MultiWriter(os.Stdout, file) to write to multiple outputs)
//	 	handlerOptions is the options for the handler, such as
//		Level is the minimum level to log
//		AddSource if true, adds source information to the log
//
// Deprecated: use NewTextSLogger instead
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
//
// Deprecated: use NewJsonSLogger instead
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
	if !s.Logger.Enabled(context.Background(), level) {
		return
	}
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
	if !s.Logger.Enabled(context.Background(), level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, logfCorrectSource, Info/Debug/Warn/Error/Fatal]
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(template, args...), pcs[0])
	_ = s.Handler().Handle(context.Background(), r)
}

func (s SLogger) With(tags ...any) Logger {
	return &SLogger{
		s.Logger.With(tags...),
	}
}
