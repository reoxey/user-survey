package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{
		log.New(os.Stdout, "[survey] ", log.Lshortfile|log.Lmsgprefix|log.LstdFlags),
	}
}
