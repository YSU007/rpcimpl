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

var defLog Interface

func SetDefLog(log Interface) {
	defLog = log
}

func Info(format string, a ...interface{}) {
	defLog.Infof(format, a...)
}

func Debug(format string, a ...interface{}) {
	defLog.Debugf(format, a...)
}

func Warn(format string, a ...interface{}) {
	defLog.Warnf(format, a...)
}

func Error(format string, a ...interface{}) {
	defLog.Errorf(format, a...)
}
