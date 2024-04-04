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

func NewZapLogger(env LogLevel) (Logger, error) {
	callerSkip := zap.AddCallerSkip(1)

	if env == Production {
		logger, err := zap.NewProduction(callerSkip)
		if err != nil {
			panic(err)
		}
		return &ZapLogger{
			logger: logger,
		}, nil
	} else if env == Development {
		logger, err := zap.NewDevelopment(callerSkip)
		if err != nil {
			panic(err)
		}
		return &ZapLogger{
			logger: logger,
		}, nil
	} else {
		panic(fmt.Sprintf("Unknown environment. Expected %s or %s. Received %s.", Development, Production, env))
	}
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
