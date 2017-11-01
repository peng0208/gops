package common

import (
	"log"
	"os"
)

var logger *log.Logger

func Logger() *log.Logger {
	if logger == nil {
		logger = NewLogger()
	}
	return logger
}

func NewLogger() *log.Logger {
	logFile := GetConfig().Server.GeneralLog
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	newlog := log.New(f,"", log.Ldate|log.Ltime)
	return newlog
}