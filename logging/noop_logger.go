package logging

type NoopLogger struct{}

var _ Logger = (*NoopLogger)(nil)

func NewNoopLogger() Logger {
	return &NoopLogger{}
}

func (l *NoopLogger) Debug(msg string, tags ...any) {}

func (l *NoopLogger) Info(msg string, tags ...any) {}

func (l *NoopLogger) Warn(msg string, tags ...any) {}

func (l *NoopLogger) Error(msg string, tags ...any) {}

func (l *NoopLogger) Fatal(msg string, tags ...any) {}

func (l *NoopLogger) Debugm(msg string, tags map[string]any) {}

func (l *NoopLogger) Infom(msg string, tags map[string]any) {}

func (l *NoopLogger) Warnm(msg string, tags map[string]any) {}

func (l *NoopLogger) Errorm(msg string, tags map[string]any) {}

func (l *NoopLogger) Fatalm(msg string, tags map[string]any) {}

func (l *NoopLogger) Debugf(template string, args ...interface{}) {}

func (l *NoopLogger) Infof(template string, args ...interface{}) {}

func (l *NoopLogger) Warnf(template string, args ...interface{}) {}

func (l *NoopLogger) Errorf(template string, args ...interface{}) {}

func (l *NoopLogger) Fatalf(template string, args ...interface{}) {}

func (l *NoopLogger) With(tags ...any) Logger {
	return l
}
