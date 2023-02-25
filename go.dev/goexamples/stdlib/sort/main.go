package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{2, 4, 5, 1}
	// 检测是否有序
	fmt.Println(sort.IntsAreSorted(ints))

	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println(sort.IntsAreSorted(ints))

	fruits := []string{"apple", "banana", "peach", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	// 直接通过sort slice
	sort.Slice(fruits, func(i, j int) bool {
		return len(fruits[i]) > len(fruits[j])
	})
	fmt.Println(fruits)
}

// ByLength 定制排序规则 需要实现接口 签名 Len Swap Less
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
