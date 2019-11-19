package sli4go

import (
	"fmt"
	"log"
	"os"
)

type StandardLogger struct {
}

var (
	calldepth = 3
)

// SetCallDepth 设置调用层
func SetCallDepth(depth int) {
	calldepth = depth
}

func (l *StandardLogger) Fatal(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))
	os.Exit(1)
}
func (l *StandardLogger) Fatalf(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintf(format, v...))
	os.Exit(1)
}
func (l *StandardLogger) Fatalln(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))
	os.Exit(1)
}
func (l *StandardLogger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	log.Output(calldepth, s)
	panic(s)
}
func (l *StandardLogger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	log.Output(calldepth, s)
	panic(s)
}
func (l *StandardLogger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	log.Output(calldepth, s)
	panic(s)
}

func (l *StandardLogger) Print(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))
}
func (l *StandardLogger) Printf(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintf(format, v...))
}
func (l *StandardLogger) Println(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))
}

func (l *StandardLogger) Trace(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))
}
func (l *StandardLogger) Tracef(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintf(format, v...))
}
func (l *StandardLogger) Traceln(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))
}
func (l *StandardLogger) Debug(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))
}
func (l *StandardLogger) Debugf(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintf(format, v...))
}
func (l *StandardLogger) Debugln(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))
}
func (l *StandardLogger) Info(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))

}
func (l *StandardLogger) Infof(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))
}
func (l *StandardLogger) Infoln(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))

}
func (l *StandardLogger) Warn(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))
}
func (l *StandardLogger) Warnf(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintf(format, v...))
}
func (l *StandardLogger) Warnln(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))

}
func (l *StandardLogger) Error(v ...interface{}) {
	log.Output(calldepth, fmt.Sprint(v...))
}
func (l *StandardLogger) Errorf(format string, v ...interface{}) {
	log.Output(calldepth, fmt.Sprintf(format, v...))
}
func (l *StandardLogger) Errorln(v ...interface{}) {
	log.Output(calldepth, fmt.Sprintln(v...))
}
