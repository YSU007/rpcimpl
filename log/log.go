package log

// Interface ----------------------------------------------------------------------------------------------------
type Interface interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

var _defLog Interface

func Info(format string, a ...interface{}) {
	_defLog.Infof(format, a...)
}

func Debug(format string, a ...interface{}) {
	_defLog.Debugf(format, a...)
}

func Warn(format string, a ...interface{}) {
	_defLog.Warnf(format, a...)
}

func Error(format string, a ...interface{}) {
	_defLog.Errorf(format, a...)
}

// Support ----------------------------------------------------------------------------------------------------
type Support byte

const (
	DefLog Support = iota
	Zap
	Logrus
)

func Init(log Support) {
	switch log {
	case Zap:
		_defLog = retZap()
	case Logrus:
		_defLog = retLogrus()
	default:

	}
}
