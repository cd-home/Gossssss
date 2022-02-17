package testsss

import (
	"testing"
)

func TestDeferred(t *testing.T) {
	t.Log("Testing deferred")
	// reverse order output
	defer t.Log("Testing deferred 1")
	defer t.Log("Testing deferred 2")
	defer t.Log("Testing deferred 3")
}


func TestDeferredPanic(t *testing.T) {
	// catch panic
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	panic("TestDeferredPanic")
}