package log

import (
	"fmt"
	"log"
)

type LogFunc = func(format string, v ...any)
type Logger struct {
	ActiveDebug bool
}

func (logger *Logger) Log(logFunc LogFunc, level string, format string, a ...any) {
	format = fmt.Sprintf("%s %s", level, format)
	logFunc(format, a...)
}

func (logger *Logger) Error(format string, a ...any) {
	logger.Log(log.Printf, "ERROR", format, a...)
}

func (logger *Logger) Info(format string, a ...any) {
	logger.Log(log.Printf, "INFO", format, a...)
}

func (logger *Logger) Fatal(format string, a ...any) {
	logger.Log(log.Fatalf, "ERROR", format, a...)
}

func (logger *Logger) Debug(format string, a ...any) {
	if logger.ActiveDebug {
		logger.Log(log.Printf, "DEBUG", format, a...)
	}
}
