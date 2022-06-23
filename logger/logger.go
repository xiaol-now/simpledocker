package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Fields logrus.Fields

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func Info(msg string, fields map[string]interface{}) {
	logger.WithFields(fields).Info(msg)
}

func Error(msg string, fields map[string]interface{}) {
	logger.WithFields(fields).Error(msg)
}

func Fatal(msg string, fields map[string]interface{}) {
	logger.WithFields(fields).Fatal(msg)
}

func Debug(msg string, fields map[string]interface{}) {
	logger.WithFields(fields).Debug(msg)
}

func PanicErr(err error) {
	logger.Panic(err)
}
