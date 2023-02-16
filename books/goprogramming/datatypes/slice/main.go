package main

import "fmt"

func main() {
	// 声明nil切片 []T, 元素可以是任何相同的类型, 无固定的长度
	var a []int
	fmt.Println(a)

	// 字面量
	var b = []int{1, 2, 3}
	fmt.Println(b)

	// 指定位置
	var c = []int{0: 2, 1: 3, 2: 4}
	fmt.Println(c)

	// 通过数组获取到切片
	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g"}
	var s1 = arr[1:4] // b c d
	fmt.Println(s1)

	// s[i:j]   0 <= i <= j <= cap(s)  左闭右开
	// i 不填默认是0 j 不填默认是 len(s), 也就是末尾
	// i <= len(s)
	var s2 = arr[3:] // d e f g
	fmt.Println(s2)

	// 切片是引用底层的数组对象
	// 修改 index=3 位置, 观察s1 s2
	arr[3] = "o"
	fmt.Println(s1) // 修改原数组, 影响到了切片
	fmt.Println(s2)

	// 修改s1 的元素观察原数组, s2
	s1[2] = "r"
	fmt.Println(arr) // 影响原数组
	fmt.Println(s1)
	fmt.Println(s2) // 影响s2

	// 注意：
	// 切片之间是无法比较的, 切片只能和nil比较
	// 无法作为map的key, slice可能在运行时被修改

	// 切片作为函数参数传递, 是值传递, 但是由于切片包含的是引用, 所以可以修改底层的数组

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
