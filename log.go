package log

type Fields map[string]interface{}

type LoggerOptions struct {
	Prefix string
	Level  string
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(fields Fields) Logger
	WithName(name string) Logger
}
