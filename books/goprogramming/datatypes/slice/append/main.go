package main

import "fmt"

func main() {
	// nil 的切片无法进行其他操作, 但是可以append
	// 当容量不够时, 会重新申请2倍的容量 (<1024时), 复制到新的空间, 然后再添加
	var s []int
	curCap := cap(s)
	for i := 0; i < 10; i++ {
		fmt.Printf("s=%v, len=%d, cap=%d\n", s, len(s), cap(s))
		// 容量够就直接添加, 不够的话, 就开辟新的内存空间, 把值复制过去, 然后再添加

		// 需要注意的是，我们在append的时候并不知道什么时候重新内存分配, 因此我们并不知道
		// 新的slice与原先的slice是否公用底层的数组(不能确定修改是否会互相影响)
		// 为了避免造成这种不确定的状态, 通常把append的结果赋值原来的切片

		s = append(s, i)
		// 是否发生了扩容
		if curCap != cap(s) {
			fmt.Println("malloc")
			curCap = cap(s)
		}
	}

	// 更新slice变量不仅对调用append函数是必要的,
	// 实际上对应任何可能导致长度、容量或底层数组变化的操作都是必要的
	// 要正确地使用slice，需要记住尽管底层数组的元素是间接访问的, 但是slice对应结构体本身的指针、长度和容量部分是直接访问的
	// 要更新这些信息需要像 s = append(s, i), 显示的赋值操作

	// append slice
	s1 := []int{99, 98, 97}
	// 支持追加多个元素
	s = append(s, 100, 101, 102)
	// 支持传递slice, 需要展开
	s = append(s, s1...)

	// 增长, 但是不引用相同数组
	// 这里没有再赋值给s
	s2 := append(s, 100)
	fmt.Println(s)
	fmt.Println(s2)
}

func Delete[v any](s []v, i int) []v {
	return append(s[:i], s[i+1:]...)
}

func Delete2[v any](s []v, i int) []v {
	// 不影响原切片的话就需要重新开辟空间存储
	var ss []v
	ss = append(ss, s[i:]...)
	ss = append(ss, s[i+1:]...)
	return ss
}
