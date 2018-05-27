package log

import (
	"sync"
)

const (
	LevelTrace   = iota //0
	LevelDebug          //1
	LevelNotice         //2
	LevelWarning        //3
	LevelFatal          //4
)

type BaseLog struct {
	init     bool
	lock     sync.Mutex
	logger   *Logger
	logLevel int
}

func (bl *BaseLog) trace(format string, v ...interface{}) {
	if bl.logLevel > LevelTrace {
		return
	}
}

func (bl *BaseLog) debug(format string, v ...interface{}) {
	if bl.logLevel > LevelDebug {
		return
	}
}

func (bl *BaseLog) notice(format string, v ...interface{}) {
	if bl.logLevel > LevelNotice {
		return
	}
}

func (bl *BaseLog) warning(format string, v ...interface{}) {
	if bl.logLevel > LevelWarning {
		return
	}
}

func (bl *BaseLog) fatal(format string, v ...interface{}) {
	if bl.logLevel > LevelFatal {
		return
	}
}

func (bl *BaseLog) requestOut(format string, v ...interface{}) {

}

func (bl *BaseLog) requestIn(format string, v ...interface{}) {

}
