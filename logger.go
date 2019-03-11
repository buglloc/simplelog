package log

import (
	"os"
	"sync"
	"time"
)

type Logger struct {
	ctx []interface{}
	mu  sync.Mutex
}

var (
	maxLvl = InfoLevel
)

func NewLogger() Logger {
	return Logger{
		ctx: []interface{}{},
	}
}

func (l Logger) Debug(msg string, ctx ...interface{}) {
	l.write(msg, DebugLevel, ctx)
}

func (l Logger) Info(msg string, ctx ...interface{}) {
	l.write(msg, InfoLevel, ctx)
}

func (l Logger) Warn(msg string, ctx ...interface{}) {
	l.write(msg, WarnLevel, ctx)
}

func (l Logger) Error(msg string, ctx ...interface{}) {
	l.write(msg, ErrorLevel, ctx)
}

func (l Logger) Crit(msg string, ctx ...interface{}) {
	l.write(msg, CritLevel, ctx)
}

func (l Logger) Child(ctx ...interface{}) Logger {
	return Logger{
		ctx: newContext(l.ctx, ctx),
	}
}

func (l *Logger) write(msg string, lvl Lvl, ctx []interface{}) {
	if lvl > maxLvl {
		return
	}

	record := Record{
		Time: time.Now(),
		Lvl:  lvl,
		Msg:  msg,
		Ctx:  newContext(l.ctx, ctx),
	}

	l.mu.Lock()
	_, _ = os.Stderr.Write(FormatRecord(record))
	l.mu.Unlock()
}
