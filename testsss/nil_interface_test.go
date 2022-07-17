package testsss

import (
	"bytes"
	"io"
	"testing"
	"unsafe"
)

type IF interface {
	Foo()
}

type S struct{}

func (s S) Foo() {}

func TypeIsNil() S {
	return S{}
}

func PointerIsNil() *S {
	var s *S
	return s
}

func EmptyInterfaceIsNil() interface{} {
	var s S
	return s
}

func EmptyInterfacePointerIsNil() interface{} {
	var s *S
	return s
}

func InterfaceIsNil() IF {
	var s S
	return s
}

func InterfacePointerIsNil() IF {
	var s *S
	return s
}

func TestNilInterface(t *testing.T) {
	// t.Log(TypeIsNil() == nil)

	t.Log(PointerIsNil() == nil)

	// interface{} is "interface{}" not nil
	t.Log(EmptyInterfaceIsNil() == nil)
	t.Log(EmptyInterfacePointerIsNil() == nil)

	// interface is "interface" not nil
	t.Log(InterfaceIsNil() == nil)
	t.Log(InterfacePointerIsNil() == nil)

	var i io.Writer
	var b *bytes.Buffer

	t.Log(i == b)

	t.Log(i == nil)
	t.Log(b)
	t.Log(unsafe.Sizeof(b))

	g := new([5]int)
	t.Log(g)
	t.Log(unsafe.Sizeof(g))

	var m map[string]string

	v, ok := m["mike"]
	if ok {
		t.Log(v)
	} else {
		t.Log("NoExist")
	}
	
	m2 := make(map[interface{}]interface{})
	m2["1"] = 1
	m2[1] = "2"

	t.Log(m2)
}
