package example_test

import (
	"fmt"
	"testing"
)

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func TestDefineMyType(t *testing.T) {
	var e example
	fmt.Printf("%+v\n", e)
}

func TestDeclareMyType(t *testing.T) {
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

func TestUnamedType(t *testing.T) {
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)
}
