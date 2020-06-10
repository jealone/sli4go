package sli4go

import (
	"log"
	"sync"
)

var (
	defaultLog *Log = stdlog()
	initOnce   sync.Once
)

type Log struct {
	initOnce      sync.Once
	printLogger   PrintLogger
	instantLogger InstantLogger
	formatLogger  FormatLogger
	lineLogger    LineLogger
	flusher       Flusher
}

func (l *Log) init() {
	l.initOnce.Do(func() {
		log.Println("sli4go initialization ...")
	})
}

// GetLogger 获取 Log 实例
func GetLogger() *Log {
	initOnce.Do(func() {
		log.Fatalln("logger must initialize for sli4go first")
	})
	return defaultLog
}

// InitLogger 初始化 Log
func InitLogger(logger interface{}) *Log {

	initOnce.Do(func() {
		l := initLogger(logger)

		l.initOnce.Do(func() {
			log.Fatalln("logger type for initialization sli4go didn't support")
		})

		defaultLog = l
	})

	return defaultLog
}

func initLogger(log interface{}) *Log {
	var l Log
	if logger, ok := log.(PrintLogger); ok {
		l.printLogger = logger
		l.init()
	}

	if logger, ok := log.(InstantLogger); ok {
		l.instantLogger = logger
		l.init()
	}

	if logger, ok := log.(FormatLogger); ok {
		l.formatLogger = logger
		l.init()
	}

	if logger, ok := log.(LineLogger); ok {
		l.lineLogger = logger
		l.init()
	}

	if logger, ok := log.(Flusher); ok {
		l.flusher = logger
		l.init()
	}

	return &l
}

func (l *Log) Fatal(v ...interface{}) {
	l.instantLogger.Fatal(v...)
}
func (l *Log) Fatalf(format string, v ...interface{}) {
	l.formatLogger.Fatalf(format, v...)
}
func (l *Log) Fatalln(v ...interface{}) {
	l.lineLogger.Fatalln(v...)
}
func (l *Log) Panic(v ...interface{}) {
	l.instantLogger.Panic(v...)
}
func (l *Log) Panicf(format string, v ...interface{}) {
	l.formatLogger.Panicf(format, v...)
}
func (l *Log) Panicln(v ...interface{}) {
	l.lineLogger.Panicln(v...)
}

func (l *Log) Print(v ...interface{}) {
	l.printLogger.Print(v...)
}
func (l *Log) Printf(format string, v ...interface{}) {
	l.printLogger.Printf(format, v...)
}
func (l *Log) Println(v ...interface{}) {
	l.printLogger.Println(v...)
}

func (l *Log) Trace(v ...interface{}) {
	l.instantLogger.Trace(v...)
}
func (l *Log) Tracef(format string, v ...interface{}) {
	l.formatLogger.Tracef(format, v...)
}
func (l *Log) Traceln(v ...interface{}) {
	l.lineLogger.Traceln(v...)
}
func (l *Log) Debug(v ...interface{}) {
	l.instantLogger.Debug(v...)
}
func (l *Log) Debugf(format string, v ...interface{}) {
	l.formatLogger.Debugf(format, v...)
}
func (l *Log) Debugln(v ...interface{}) {
	l.lineLogger.Debugln(v...)
}
func (l *Log) Info(v ...interface{}) {
	l.instantLogger.Info(v...)
}
func (l *Log) Infof(format string, v ...interface{}) {
	l.formatLogger.Infof(format, v...)
}
func (l *Log) Infoln(v ...interface{}) {
	l.lineLogger.Infoln(v...)
}
func (l *Log) Warn(v ...interface{}) {
	l.instantLogger.Warn(v...)
}
func (l *Log) Warnf(format string, v ...interface{}) {
	l.formatLogger.Warnf(format, v...)
}
func (l *Log) Warnln(v ...interface{}) {
	l.lineLogger.Warnln(v...)
}
func (l *Log) Error(v ...interface{}) {
	l.instantLogger.Error(v...)
}
func (l *Log) Errorf(format string, v ...interface{}) {
	l.formatLogger.Errorf(format, v...)
}
func (l *Log) Errorln(v ...interface{}) {
	l.lineLogger.Errorln(v...)
}

func (l *Log) Flush() error {
	return l.flusher.Flush()
}
