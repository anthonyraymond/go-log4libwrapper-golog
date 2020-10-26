package log4libwrapper

import (
	"github.com/anthonyraymond/go-log4lib"
	"github.com/kataras/golog"
)

type gologLoggerWrapper struct {
	delegate *golog.Logger
}

func (w *gologLoggerWrapper) Info(args ...interface{}) {
	w.delegate.Info(args...)
}

func (w *gologLoggerWrapper) Warn(args ...interface{}) {
	w.delegate.Warn(args...)
}

func (w *gologLoggerWrapper) Error(args ...interface{}) {
	w.delegate.Error(args...)
}

func (w *gologLoggerWrapper) Panic(args ...interface{}) {
	w.delegate.Error(args...)
	panic(args)
}

func (w *gologLoggerWrapper) Fatal(args ...interface{}) {
	w.delegate.Fatal(args...)
}


func (w *gologLoggerWrapper) Infof(template string, args ...interface{}) {
	w.delegate.Infof(template, args...)
}

func (w *gologLoggerWrapper) Warnf(template string, args ...interface{}) {
	w.delegate.Warnf(template, args...)
}

func (w *gologLoggerWrapper) Errorf(template string, args ...interface{}) {
	w.delegate.Errorf(template, args...)
}

func (w *gologLoggerWrapper) Panicf(template string, args ...interface{}) {
	w.delegate.Errorf(template, args...)
	panic(args)
}

func (w *gologLoggerWrapper) Fatalf(template string, args ...interface{}) {
	w.delegate.Fatalf(template, args...)
}

func WrapGologLogger(pointerToLogger *golog.Logger) log4lib.LibLogger {
	return &gologLoggerWrapper{delegate: pointerToLogger}
}
