package x

import (
	"fmt"
	"sync"
	"testing"
)

type Iface interface {
	Foo()
}

type A struct{}

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

type Book struct {
	Pages int
}

func TestStructAddress(t *testing.T) {
	var book = Book{}
	t.Log(&book)

	t.Log(&Book{})

	// 不可寻址, 这里的 &Book{} 字面量 并没有地址
	// &Book{}.Pages

	// 可寻址
	b := &Book{}
	b.Pages = 1
	t.Log(*b)
}

type Data struct {
	number int
	sync.Mutex
}

// dp 不能是指针类型
type dp Data

func (d dp) Add() {
	defer d.Unlock()
	d.Lock()
	d.number++
}

func TestMutxt(t *testing.T) {

}
