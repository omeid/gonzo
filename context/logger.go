package context

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

var std *logrus.Entry

func init() {

	fmt.Println(os.Getenv("GONZO_LOG"))
	l := logrus.New()
	switch os.Getenv("GONZO_LOG") {
	case "debug":
		l.Level = logrus.DebugLevel
	case "info":
		l.Level = logrus.InfoLevel
	case "warn":
		l.Level = logrus.WarnLevel
	case "error":
		l.Level = logrus.ErrorLevel
	case "fatal":
		l.Level = logrus.FatalLevel
	case "panic":
		l.Level = logrus.PanicLevel
	default:
		l.Level = logrus.InfoLevel
	}
	std = logrus.NewEntry(l)
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
}
