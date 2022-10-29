package main

import "fmt"

func main() {
	// 每个元素被初始化为 元素类型的零值
	var arr [3]int
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 字面量
	arr1 := [3]int{1, 2, 3}
	// 指定位置, 未指定位置的默认是零值
	arr2 := [3]int{0: 1, 2: 3}
	fmt.Println(arr2)
	fmt.Println(arr1)

	// 自动计算长度
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	// 1. 如果数组的元素是可比较的, 那么数组也是可比较的
	// 2. 函数是值传递(复制), 传递数组, 函数内部修改数组不会影响原数组
	// 如果想要修改, 可传递数组指针, 或者使用切片
}
