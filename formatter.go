package log

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	timeFormat     = "2006-01-02T15:04:05-0700"
	termTimeFormat = "15:04:05"
	termMsgJust    = 40
)

var (
	colored = terminal.IsTerminal(int(os.Stderr.Fd()))
)

func FormatRecord(r Record) []byte {
	var color = 0
	if colored {
		switch r.Lvl {
		case CritLevel:
			color = 35
		case ErrorLevel:
			color = 31
		case WarnLevel:
			color = 33
		case InfoLevel:
			color = 32
		case DebugLevel:
			color = 36
		}
	}

	b := &bytes.Buffer{}
	if color > 0 {
		_, _ = fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s] %s ", color, r.Lvl, r.Time.Format(termTimeFormat), r.Msg)
	} else {
		_, _ = fmt.Fprintf(b, "[%s] [%s] %s ", r.Lvl, r.Time.Format(timeFormat), r.Msg)
	}

	// try to justify the log output for short messages
	if len(r.Ctx) > 0 && len(r.Msg) < termMsgJust {
		b.Write(bytes.Repeat([]byte{' '}, termMsgJust-len(r.Msg)))
	}

	// print the keys logfmt style
	logfmt(b, r.Ctx, color)
	return b.Bytes()
}

func logfmt(buf *bytes.Buffer, ctx []interface{}, color int) {
	for i := 0; i < len(ctx); i += 2 {
		if i != 0 {
			buf.WriteByte(' ')
		}

		k, ok := ctx[i].(string)
		v := formatLogfmtValue(ctx[i+1])
		if !ok {
			k, v = errorKey, formatLogfmtValue(ctx[i])
		}

		// XXX: we should probably check that all of your key bytes aren't invalid
		if color > 0 {
			_, _ = fmt.Fprintf(buf, "\x1b[%dm%s\x1b[0m=%s", color, k, v)
		} else {
			_, _ = fmt.Fprintf(buf, "%s=%s", k, v)
		}
	}

	buf.WriteByte('\n')
}

func formatLogfmtValue(value interface{}) string {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	return fmt.Sprintf("%q", stringVal)
}
