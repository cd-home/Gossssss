package main

import "fmt"

// 范型
// 类型参数

// MapKeys 类型参数 type parameters.
// comparable 是可比较的 any 表示任意类型 interface{}
func MapKeys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

// List type struct
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val  T
	next *element[T]
}

// Push
// 类型参数通用可以作用与方法, 前提是接收者类型定义类型
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	m := map[int]string{1: "1", 2: "2", 3: "3"}
	// 可以写类型
	fmt.Println(MapKeys[int, string](m))
	// 或者编译器自动推断
	fmt.Println(MapKeys(m))

	lst := List[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)
	fmt.Println(lst.GetAll())
}
