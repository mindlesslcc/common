package logs

const (
	ERROR = iota
	WARN
	DEBUG
	INFO
)

type loggerType func() Logger

type Logger interface {
	Init(string) error
	WriteMsg(level int, msg string) error
	Destroy()
	Flush()
}

var loggers = make(map[string]loggerType)
var logger Logger

func Register(name string, log loggerType) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, dup := loggers[name]; dup {
		panic("logs: Register called twice for provider " + name)
	}
	loggers[name] = log
}

func SetLogger(loggerName, args string) (err error) {
	if _, exist := loggers[loggerName]; exist != true {
		panic("logs: set logger which don't exist")
	}
	NewLoggerFunc := loggers[loggerName]
	logger = NewLoggerFunc()
	err = logger.Init(args)
	return err
}

func Info(msg string) error {
	return logger.WriteMsg(INFO, msg)
}

func Debug(msg string) error {
	return logger.WriteMsg(DEBUG, msg)
}

func Warn(msg string) error {
	return logger.WriteMsg(WARN, msg)
}

func Error(msg string) error {
	return logger.WriteMsg(ERROR, msg)
}
