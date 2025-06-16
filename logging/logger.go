package logging

import grpclogging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

type Logger interface {
	Debug(msg string, tags ...any)

	Info(msg string, tags ...any)

	Warn(msg string, tags ...any)

	Error(msg string, tags ...any)

	Fatal(msg string, tags ...any)

	Debugf(template string, args ...interface{})

	Infof(template string, args ...interface{})

	Warnf(template string, args ...interface{})

	Errorf(template string, args ...interface{})

	Fatalf(template string, args ...interface{})

	With(tags ...any) Logger

	InterceptorLogger() grpclogging.Logger
}
