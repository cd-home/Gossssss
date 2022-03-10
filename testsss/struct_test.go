package testsss

import (
	"fmt"
	"testing"
)

type Iface interface {
	Foo()
}

type A struct {}
func (a A) Foo() {
	fmt.Println("A.Foo")
}

func (a A) Bar() {
	fmt.Println("A.Bar")
}


type B struct {
	A
}

func (b B) Foo() {
	fmt.Println("b.Foo")
}


func printType(i Iface) {
	if a, ok := i.(*A); ok {
		a.Foo()
		a.Bar()
	}
	if b, ok := i.(*B); ok {
		b.Foo()
		b.Bar()
	}
}

func TestStruct(t *testing.T) {
	t.Helper()
	printType(&A{})
	printType(&B{})
}