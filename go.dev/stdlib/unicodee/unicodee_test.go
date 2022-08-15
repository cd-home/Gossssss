package stdsss

import (
	"testing"
	"unicode"
)

func TestUnicode(t *testing.T) {
	s := "123abc你"
	for pos, char := range s {
		if unicode.IsDigit(char) {
			t.Logf("pos = %d, char = %d, encode is digit", pos, char)
		}
		if unicode.IsLetter(char) {
			t.Logf("pos = %d, char = %d, encode is a letter", pos, char)
		}
	}
}
