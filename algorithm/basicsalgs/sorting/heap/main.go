package main

import (
	"fmt"
)

// 保证父节点大于两个子节点
func sink(tree []int, n, i int) {
	if i >= n {
		return
	}
	c1 := 2*i + 1 // 左子节点
	c2 := 2*i + 2 // 右子节点

	max := i // 默认最大节点是当前节点
	if c1 < n && tree[c1] > tree[max] {
		// 最大节点是左节点
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		// 最大节点是右节点
		max = c2
	}

	// 如果最大节点不是当前节点
	if max != i {
		// 交换最大节点和子节点
		tree[max], tree[i] = tree[i], tree[max]
		// 对子节点重新堆化
		sink(tree, n, max)
	}
}

// 数组建堆
func buildHeap(tree []int, n int) {
	lastNode := n - 1            // 最后一个节点
	parent := (lastNode - 1) / 2 // 最后一个节点的父节点
	for i := parent; i >= 0; i-- {
		sink(tree, n, i)
	}
}

func heapSort(tree []int, n int) {
	buildHeap(tree, n) // 建堆
	fmt.Println(tree)
	for i := n - 1; i >= 0; i-- {
		// 0 为堆顶元素  必定是最大
		tree[i], tree[0] = tree[0], tree[i]
		sink(tree, i, 0) // 继续堆化
	}
}

func main() {
	arr := []int{10, 6, 2, 1, 12, 9, 0, 7, 6}
	heapSort(arr, len(arr))
	fmt.Println(arr)
}
