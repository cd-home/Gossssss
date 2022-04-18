package main

import (
	"container/heap"
	"fmt"
)

func main() {
	myheap := &heapFloat32{1.2, 2.1, 3.1, -100.1}
	heap.Init(myheap)
	size := len(*myheap)
	fmt.Println(size)
}

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old) - 1]
	new_ := old[0:len(old) - 1]
	*n = new_
	return x
}

func (n *heapFloat32) Push(x interface{})  {
	*n = append(*n, x.(float32))
}

func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}