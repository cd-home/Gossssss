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

	t.Log(*ps)       // 12345
	t.Log(string(b)) // 103
}
func TestStringSlice(t *testing.T) {
	s := "123"
	b := s[0]

	// b is a byte 
	t.Logf("%v, %v, %T", b, string(b), b) // 49, 1, uint8

	// v is a byte 
	for _, v := range s {
		t.Logf("%v, %v, %T", v, string(b), v) // 49, 1, int32   50, 1, int32   51, 1, int32
	}
}
