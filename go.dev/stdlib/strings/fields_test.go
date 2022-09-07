package strings

import (
	"strings"
	"testing"
)

func TestStringsFields(t *testing.T) {
	line := "How do you do ?"
	fields := strings.Fields(line)
	t.Log(fields)
}
