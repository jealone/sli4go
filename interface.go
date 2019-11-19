package sli4go

type LogInterface interface {
	InstantLog
	FormatLog
	LineLog
}

type PrintLog interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

type InstantLog interface {
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
}

type FormatLog interface {
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Panicf(string, ...interface{})
}

type LineLog interface {
	Traceln(...interface{})
	Debugln(...interface{})
	Infoln(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
	Fatalln(...interface{})
	Panicln(...interface{})
}
