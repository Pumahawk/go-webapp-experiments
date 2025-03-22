package log

import (
	"fmt"
	"log"
)

type Logger struct {
	ActiveDebug bool
}

func (logger *Logger) Log(level string, format string, a ...any) {
	format = fmt.Sprintf("%s %s", level, format)
	log.Printf(format, a...)
}

func (logger *Logger) Error(format string, a ...any) {
	logger.Log("ERROR", format, a...)
}

func (logger *Logger) Info(format string, a ...any) {
	logger.Log("INFO", format, a...)
}

func (logger *Logger) Debug(format string, a ...any) {
	if logger.ActiveDebug {
		logger.Log("DEBUG", format, a...)
	}
}
