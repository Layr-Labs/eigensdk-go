package logging

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
}
