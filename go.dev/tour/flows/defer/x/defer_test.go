package x

import "testing"

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
