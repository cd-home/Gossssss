package testsss

import (
	"testing"
)

func TestBytes(t *testing.T) {
	s := "123"
	ps := &s

	// Copy to b
	b := []byte(*ps)

	s += "4"
	*ps += "5"

	b[1] = '0'

	t.Log(*ps)
	t.Log(string(b))
}