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
}
