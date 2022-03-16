package snippetsss

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

func TestIota(t *testing.T) {
	t.Log(KB)
	t.Log(MB)
	t.Log(GB)
	t.Log(TB)
}

func TestFormatTime(t *testing.T) {
	t.Log(time.Now().Format(FormatTime2))
}
