package strings_test

import (
	"strings"
	"testing"
)

func TestStringsbuilder(t *testing.T) {
	var buf strings.Builder
	buf.Grow(10)
	buf.WriteString("hello world")
	t.Log(buf.String())

	buf.WriteString(" AND A = 1")
	buf.WriteString(" OR B = 2")
	buf.WriteString(" AND C = 3")

	t.Log(buf.String())
}
