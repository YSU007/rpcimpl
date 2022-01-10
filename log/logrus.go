package log

import (
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)

	log.Info()
}

type Logrus struct{}

func (l Logrus) Debug(v ...interface{}) {
	log.Debug(v...)
}

func (l Logrus) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func (l Logrus) Info(v ...interface{}) {
	log.Info(v...)
}

func (l Logrus) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func (l Logrus) Warn(v ...interface{}) {
	log.Warn(v...)
}

func (l Logrus) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func (l Logrus) Error(v ...interface{}) {
	log.Error(v...)
}

func (l Logrus) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func (l Logrus) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (l Logrus) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (l Logrus) Panic(v ...interface{}) {
	log.Panic(v...)
}

func (l Logrus) Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
