package _defer

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

type Key struct {
	Value []int
}

func (k *Key) Add(v int) *Key {
	k.Value = append(k.Value, v)
	return k
}

func DeferDelayCall() *Key {
	k := &Key{}

	// first call add(1) to simple method, defer k.Add(2)
	defer k.Add(1).Add(2)

	k.Add(3)

	return k
}

func TestDeferDelayCall(t *testing.T) {
	k := DeferDelayCall()
	t.Log(k.Value)
}
