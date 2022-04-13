package logging

import (
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init()  {
	logger = logrus.New()
	logger.SetFormatter(
		&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	logger.SetLevel(logrus.TraceLevel)
}

func Info(args ...interface{}){
	logger.Info(args...)
}


func Warn(args ...interface{}){
	logger.Warn(args...)
}

func Debug(args ...interface{}){
	logger.Debug(args...)
}

func Error(args ...interface{}){
	logger.Error(args...)
}

func Panic(args ...interface{}){
	logger.Panic(args...)
}

func WithFields(args logrus.Fields) *logrus.Entry{
	return logger.WithFields(args)
}