package logger

import (
	"io"
	"log"
	"os"
)

type LoggingLevel uint8

const (
	ERROR LoggingLevel = iota
	WARNING
	INFO
	NONE
)

type Logger struct {
	prefix string
	level  LoggingLevel
	log    *log.Logger
	Out    io.Writer
}

func New(prefix string, level LoggingLevel) *Logger {
	logger := &Logger{
		prefix: prefix,
		level:  level,
		log:    log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile),
	}
	logger.Out = logger.log.Writer()
	return logger
}

func (logger Logger) Info(message ...any) {
	if logger.level > INFO {
		return
	}

	logger.log.Println(message...)
}
func (logger Logger) Error(message ...any) {
	if logger.level > ERROR {
		return
	}

	logger.log.Println(message...)
}
func (logger Logger) Warning(message ...any) {
	if logger.level > WARNING {
		return
	}

	logger.log.Println(message...)
}

// implements io.Writer interface
func (logger Logger) Write(p []byte) (n int, err error) {
	logger.log.Println(string(p[:]))
	return len(p), nil
}
