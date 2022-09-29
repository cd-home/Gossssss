package consts

import (
	"testing"
	"time"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

const FormatTime1 = "2006-01-02 15:04:05"
const FormatTime2 = "2006/01/02 15:04:05"

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel int8 = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	// _minLevel = DebugLevel
	// _maxLevel = FatalLevel
)

func TestIota(t *testing.T) {
	t.Log(KB)
	t.Log(MB)
	t.Log(GB)
	t.Log(TB)
}

func TestFormatTime(t *testing.T) {
	t.Log(time.Now().Format(FormatTime2))
}

func TestLogMode(t *testing.T) {
	t.Log(DebugLevel)
	t.Log(InfoLevel)
	t.Log(WarnLevel)
	t.Log(ErrorLevel)
}
