package logger

import (
	"log"
	"strings"
)

var levelsMap = map[string]int{
	"ERROR": 2,
	"INFO":  5,
	"DEBUG": 8,
}

func InitLogger(level string) *Logger {
	logger := Logger{}
	logger.LEVEL = levelsMap[strings.ToUpper(level)]
	return &logger
}

type Logger struct {
	LEVEL int
}

func (l *Logger) Info(msg string) {
	if l.LEVEL > 4 {
		log.Println("[INFO]", msg)
	}
}
