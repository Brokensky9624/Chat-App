package utils

import (
	"log"
	"os"
)

const (
	appName = "[Chat App]"
)

var (
	Logger          *log.Logger
	LoggerCallDepth = 2
)

func Output(s string) {
	Logger.Output(LoggerCallDepth, s)
}

func init() {
	Logger = log.New(os.Stdout, appName, log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}
