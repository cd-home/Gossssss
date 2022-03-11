package testsss

import (
	"testing"
)


func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	panic("TestDeferredPanic")
}