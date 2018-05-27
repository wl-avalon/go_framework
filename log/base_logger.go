package log

import (
	"time"
)

type Logger interface {
	Init(config string) error
	WriteLog(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}
