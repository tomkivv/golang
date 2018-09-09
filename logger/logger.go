package util

import (
	"sync"
)

type logger struct {
}

var instance *logger
var once sync.Once

func GetInstance() *logger {
	once.Do(func() {
		instance = &logger{}
	})
	return instance
}
