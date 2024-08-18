package logger

import (
	"log"
	"os"
)

// Logger is a custom logger for the application.
type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

// NewLogger creates a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// InfoMsg Info logs informational messages.
func (l *Logger) InfoMsg(msg string) {
	l.infoLog.Println(msg)
}

// ErrorMsg Error logs error messages.
func (l *Logger) ErrorMsg(err error) {
	l.errorLog.Println(err)
}
