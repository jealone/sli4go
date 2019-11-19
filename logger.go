package sli4go

import (
	"sync"
)

var defaultLogger = initLogger(&StandardLogger{})

type Logger struct {
	initOnce      sync.Once
	printLogger   PrintLog
	instantLogger InstantLog
	formatLogger  FormatLog
	lineLogger    LineLog
}

func (l *Logger) init() {
	l.initOnce.Do(func() {
		defaultLogger.Infoln("sli4go initialization ...")
	})
}

// GetLogger 获取 Logger 实例
func GetLogger() *Logger {
	return defaultLogger
}

// InitLogger 初始化 Logger
func InitLogger(logger interface{}) *Logger {

	l := initLogger(logger)

	l.initOnce.Do(func() {
		defaultLogger.Fatalln("logger type for initialization sli4go didn't support")
	})

	defaultLogger = l

	return l
}

func initLogger(log interface{}) *Logger {
	var l Logger
	if logger, ok := log.(PrintLog); ok {
		l.printLogger = logger
		l.init()
	}

	if logger, ok := log.(InstantLog); ok {
		l.instantLogger = logger
		l.init()
	}

	if logger, ok := log.(FormatLog); ok {
		l.formatLogger = logger
		l.init()
	}

	if logger, ok := log.(LineLog); ok {
		l.lineLogger = logger
		l.init()
	}

	return &l
}

func (l *Logger) Fatal(v ...interface{}) {
	l.instantLogger.Fatal(v...)
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.formatLogger.Fatalf(format, v...)
}
func (l *Logger) Fatalln(v ...interface{}) {
	l.lineLogger.Fatalln(v...)
}
func (l *Logger) Panic(v ...interface{}) {
	l.instantLogger.Panic(v...)
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.formatLogger.Panicf(format, v...)
}
func (l *Logger) Panicln(v ...interface{}) {
	l.lineLogger.Panicln(v...)
}

func (l *Logger) Print(v ...interface{}) {
	l.printLogger.Print(v...)
}
func (l *Logger) Printf(format string, v ...interface{}) {
	l.printLogger.Printf(format, v...)
}
func (l *Logger) Println(v ...interface{}) {
	l.printLogger.Println(v...)
}

func (l *Logger) Trace(v ...interface{}) {
	l.instantLogger.Trace(v...)
}
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.formatLogger.Tracef(format, v...)
}
func (l *Logger) Traceln(v ...interface{}) {
	l.lineLogger.Traceln(v...)
}
func (l *Logger) Debug(v ...interface{}) {
	l.instantLogger.Debug(v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.formatLogger.Debugf(format, v...)
}
func (l *Logger) Debugln(v ...interface{}) {
	l.lineLogger.Debugln(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.instantLogger.Info(v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.formatLogger.Infof(format, v...)
}
func (l *Logger) Infoln(v ...interface{}) {
	l.lineLogger.Infoln(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.instantLogger.Warn(v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.formatLogger.Warnf(format, v...)
}
func (l *Logger) Warnln(v ...interface{}) {
	l.lineLogger.Warnln(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.instantLogger.Error(v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.formatLogger.Errorf(format, v...)
}
func (l *Logger) Errorln(v ...interface{}) {
	l.lineLogger.Errorln(v...)
}
