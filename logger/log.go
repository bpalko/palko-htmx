package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func GetLogger() *logrus.Logger {
	return log
}

func SetOutput(output *os.File) {
	log.SetOutput(output)
}

func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}
