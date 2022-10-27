package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	//fmt.Println(Remove(s, 2))
	fmt.Println(Delete(s, 2))
}

// Remove
// 删除元素
func Remove[v any](s []v, i int) []v {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func Delete[v any](s []v, i int) []v {
	return append(s[:i], s[i+1:]...)
}
