package sli4go

func Fatal(v ...interface{}) {
	defaultLog.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	defaultLog.Fatalf(format, v...)
}
func Fatalln(v ...interface{}) {
	defaultLog.Fatalln(v...)
}
func Panic(v ...interface{}) {
	defaultLog.Panic(v...)
}
func Panicf(format string, v ...interface{}) {
	defaultLog.Panicf(format, v...)
}
func Panicln(v ...interface{}) {
	defaultLog.Panicln(v...)
}

func Print(v ...interface{}) {
	defaultLog.Print(v...)
}
func Printf(format string, v ...interface{}) {
	defaultLog.Printf(format, v...)
}
func Println(v ...interface{}) {
	defaultLog.Println(v...)
}

func Trace(v ...interface{}) {
	defaultLog.Trace(v...)
}
func Tracef(format string, v ...interface{}) {
	defaultLog.Tracef(format, v...)
}
func Traceln(v ...interface{}) {
	defaultLog.Traceln(v...)
}
func Debug(v ...interface{}) {
	defaultLog.Debug(v...)
}
func Debugf(format string, v ...interface{}) {
	defaultLog.Debugf(format, v...)
}
func Debugln(v ...interface{}) {
	defaultLog.Debugln(v...)
}
func Info(v ...interface{}) {
	defaultLog.Info(v...)
}
func Infof(format string, v ...interface{}) {
	defaultLog.Infof(format, v...)
}
func Infoln(v ...interface{}) {
	defaultLog.Infoln(v...)
}
func Warn(v ...interface{}) {
	defaultLog.Warn(v...)
}
func Warnf(format string, v ...interface{}) {
	defaultLog.Warnf(format, v...)
}
func Warnln(v ...interface{}) {
	defaultLog.Warnln(v...)
}
func Error(v ...interface{}) {
	defaultLog.Error(v...)
}
func Errorf(format string, v ...interface{}) {
	defaultLog.Errorf(format, v...)
}
func Errorln(v ...interface{}) {
	defaultLog.Errorln(v...)
}

func Flush() error {
	return defaultLog.Flush()
}
