package testsss

import (
	"testing"
)

func TestSwitches(t *testing.T) {
	a, b, c, d := 1, 2, 3, 4
	number := 4
	// switch-case no need to break
	switch number {
	case a:
		t.Log("a == number")
	case b:
		t.Log("b == number")
	case c:
		t.Log("c == number")
	case d:
		t.Log("d == number")
	}
}