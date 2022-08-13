package testsss

import (
	"testing"
)

func TestUTF8(t *testing.T) {
	t.Log(len("A"))
	t.Log(len("你"))
	t.Log(len("你好"))
}

func TestRangeStr(t *testing.T) {
	s := "123abc你"
	for _, v := range s {
		t.Logf("rune %v", v)
		t.Log(string(v))

	}
}
