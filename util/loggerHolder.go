package util

import (
	"github.com/sirupsen/logrus"
	"sync"
)


var instance *logrus.Logger
var once sync.Once

func GetLoggerInstance() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()
		instance.Level = logrus.InfoLevel
		instance.Formatter = &logrus.JSONFormatter{}

	})
	return instance
}


