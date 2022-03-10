package testsss

import (
	"testing"
)

func TestBytes(t *testing.T) {
	s := "123"
	ps := &s

	// Copy to b, A new space
	b := []byte(*ps)

	s += "4"
	*ps += "5"

	b[1] = '0'

	t.Log(*ps) // 12345
	t.Log(string(b)) // 103
}