package sli4go

import (
	"fmt"
	"io"
	"log"
	"os"
)

type StandardLogger struct {
}

func (l *StandardLogger) Writer() io.Writer {
	return log.Writer()
}

func (l *StandardLogger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}
func (l *StandardLogger) Fatalf(format string, v ...interface{}) {
	log.Output(3, fmt.Sprintf(format, v...))
	os.Exit(1)
	//log.Fatalf(format, v...)
}
func (l *StandardLogger) Fatalln(v ...interface{}) {
	log.Output(3, fmt.Sprintln(v...))
	os.Exit(1)
	//log.Fatalln(v...)
}
func (l *StandardLogger) Panic(v ...interface{}) {
	log.Panic(v...)
}
func (l *StandardLogger) Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
func (l *StandardLogger) Panicln(v ...interface{}) {
	log.Panicln(v...)
}

func (l *StandardLogger) Print(v ...interface{}) {
	log.Print(v...)
}
func (l *StandardLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *StandardLogger) Println(v ...interface{}) {
	log.Println(v...)
}

func (l *StandardLogger) Trace(v ...interface{}) {
	log.Print(v...)
}
func (l *StandardLogger) Tracef(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *StandardLogger) Traceln(v ...interface{}) {
	log.Println(v...)
}
func (l *StandardLogger) Debug(v ...interface{}) {
	log.Print(v...)
}
func (l *StandardLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *StandardLogger) Debugln(v ...interface{}) {
	log.Println(v...)
}
func (l *StandardLogger) Info(v ...interface{}) {
	log.Print(v...)
}
func (l *StandardLogger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *StandardLogger) Infoln(v ...interface{}) {
	log.Println(v...)
}
func (l *StandardLogger) Warn(v ...interface{}) {
	log.Print(v...)
}
func (l *StandardLogger) Warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *StandardLogger) Warnln(v ...interface{}) {
	log.Println(v...)
}
func (l *StandardLogger) Error(v ...interface{}) {
	log.Print(v...)
}
func (l *StandardLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *StandardLogger) Errorln(v ...interface{}) {
	log.Println(v...)
}
