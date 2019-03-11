package log

import (
	"time"
)

type (
	Lvl int

	Record struct {
		Time time.Time
		Lvl  Lvl
		Msg  string
		Ctx  []interface{}
	}
)

const (
	CritLevel Lvl = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel

	errorKey = "LOG_ERROR"
)

// Returns the name of a Lvl
func (l Lvl) String() string {
	switch l {
	case DebugLevel:
		return "DBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "EROR"
	case CritLevel:
		return "CRIT"
	default:
		panic("bad level")
	}
}

func newContext(prefix []interface{}, suffix []interface{}) []interface{} {
	normalizedSuffix := normalize(suffix)
	newCtx := make([]interface{}, len(prefix)+len(normalizedSuffix))
	n := copy(newCtx, prefix)
	copy(newCtx[n:], normalizedSuffix)
	return newCtx
}

func normalize(ctx []interface{}) []interface{} {
	if len(ctx)%2 != 0 {
		ctx = append(ctx, nil, errorKey, "normalized odd number of arguments")
	}
	return ctx
}
