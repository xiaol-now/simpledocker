package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Fields logrus.Fields

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
