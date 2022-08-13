package testsss

import (
	"bytes"
	"testing"
)

func TestBuffer(t *testing.T) {
	
	var cache bytes.Buffer 
	// cache := new(bytes.Buffer)

	cache.Write([]byte("Value"))

	t.Log(cache.String())


	// Read from cache
	buf := make([]byte, 10)
	cache.Read(buf)

	t.Log(string(buf))
}