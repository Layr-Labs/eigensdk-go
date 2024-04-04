package logging

import (
	"fmt"

	"go.uber.org/zap"
)

type LogLevel string

const (
	Development LogLevel = "development" // prints debug and above
	Production  LogLevel = "production"  // prints info and above
)

type ZapLogger struct {
	logger *zap.Logger
}

var _ Logger = (*ZapLogger)(nil)

// NewZapLogger creates a new logger wrapped the zap.Logger
func NewZapLogger(env LogLevel) (Logger, error) {
	var config zap.Config

	if env == Production {
		config = zap.NewProductionConfig()
	} else if env == Development {
		config = zap.NewDevelopmentConfig()
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}

	return NewZapLoggerByConfig(config, zap.AddCallerSkip(1))
}

// NewZapLoggerByConfig creates a logger wrapped the zap.Logger
// Note if the logger need to show the caller, need use `zap.AddCallerSkip(1)` ad options
func NewZapLoggerByConfig(config zap.Config, options ...zap.Option) (Logger, error) {
	logger, err := config.Build(options...)
	if err != nil {
		panic(err)
	}

	return &ZapLogger{
		logger: logger,
	}, nil
}

func (z *ZapLogger) Debug(msg string, tags ...any) {
	z.logger.Sugar().Debugw(msg, tags...)
}

func (z *ZapLogger) Info(msg string, tags ...any) {
	z.logger.Sugar().Infow(msg, tags...)
}

func (z *ZapLogger) Warn(msg string, tags ...any) {
	z.logger.Sugar().Warnw(msg, tags...)
}

func (z *ZapLogger) Error(msg string, tags ...any) {
	z.logger.Sugar().Errorw(msg, tags...)
}

func (z *ZapLogger) Fatal(msg string, tags ...any) {
	z.logger.Sugar().Fatalw(msg, tags...)
}

func (z *ZapLogger) Debugf(template string, args ...interface{}) {
	z.logger.Sugar().Debugf(template, args...)
}

func (z *ZapLogger) Infof(template string, args ...interface{}) {
	z.logger.Sugar().Infof(template, args...)
}

func (z *ZapLogger) Warnf(template string, args ...interface{}) {
	z.logger.Sugar().Warnf(template, args...)
}

func (z *ZapLogger) Errorf(template string, args ...interface{}) {
	z.logger.Sugar().Errorf(template, args...)
}

func (z *ZapLogger) Fatalf(template string, args ...interface{}) {
	z.logger.Sugar().Fatalf(template, args...)
}

func (z *ZapLogger) With(tags ...any) Logger {
	return &ZapLogger{
		logger: z.logger.Sugar().With(tags...).Desugar(),
	}
}
