package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	// 反转切片
	Reverse(s)
	fmt.Println(s)

	fmt.Println(Remove(s, 2))

	ss := []string{"a", "b", "", "d", "e", ""}
	fmt.Println(nonEmpty(ss))
	fmt.Println(ss)
}

// Remove
// 删除元素, 保持原有的顺序
func Remove[v any](s []v, i int) []v {
	// s = 1 2 3 4 5 6 7 8
	//         3 4 5 6 7 8
	//         4 5 6 7 8
	copy(s[i:], s[i+1:])
	// 会多最后一个, 干掉最后1个
	return s[:len(s)-1]
}

// Remove2
// 如果不需要保持有序, 用末尾元素覆盖需要删除的位置的元素
func Remove2[v any](s []v, i int) []v {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Reverse[v any](s []v) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// nonEmpty
// 在原有的内存基础上, 返回非空的字符串序列, 避免了另外分配内存
func nonEmpty(ss []string) []string {
	// ss := ["a", "b", "", "d", "e"]
	i := 0
	// 原地修改
	for _, s := range ss {
		if s != "" {
			ss[i] = s
			i++
		}
	}
	// 输出的与ss共享底层数组
	return ss[:i]
}

func nonEmpty2(ss []string) []string {
	// 利用原切片的底层数组
	out := ss[:0]
	for _, s := range ss {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// Push
// 模拟栈操作
func Push[v any](stack []v, item v) []v {
	return append(stack, item)
}

func Pop[v any](stack []v) v {
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return top
}
