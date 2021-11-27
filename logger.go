package glog

import (
	"fmt"
	"log"
	"os"
)

type Env string

const (
	Dev  Env = "dev"
	Prod Env = "prod"
)

// Logger - provides convenient interface for logging
type Logger interface {
	Error(err error)
	Debugf(format string, args ...interface{})
	Noticef(format string, args ...interface{})
}

func NewStdoutLogger(env Env, namespace string) *StdLogger {
	debugLogger := log.New(os.Stdout, ""+namespace+" : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	errLogger := log.New(os.Stderr, ""+namespace+" : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	return &StdLogger{
		debugLogger: debugLogger,
		errLogger:   errLogger,
		env:         env,
	}
}

// StdLogger - stderr & stdout
type StdLogger struct {
	errLogger   *log.Logger
	debugLogger *log.Logger
	env         Env
}

// Error - prints error message
func (l *StdLogger) Error(err error) {
	if err == nil {
		return
	}

	_ = l.errLogger.Output(2, fmt.Sprintf("ERROR : %s", err.Error()))
}

// Debugf - prints debug message only in dev mode
func (l *StdLogger) Debugf(format string, args ...interface{}) {
	if l.env == Prod {
		return
	}

	_ = l.debugLogger.Output(2, fmt.Sprintf("DEBUG : "+format, args...))
}

// Noticef - prints notice message in dev and in prod
func (l *StdLogger) Noticef(format string, args ...interface{}) {
	_ = l.debugLogger.Output(2, fmt.Sprintf("NOTICE : "+format, args...))
}
