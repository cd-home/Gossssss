package main

import "fmt"

func main() {
	// 切片不需要指定长度, 可以动态增长
	// nil slice
	var s []int
	fmt.Println(s)

	// 会初始化
	s1 := make([]int, 3, 3)
	fmt.Println(s1)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	fmt.Println(s1)

	// 指定 len=cap=3
	s2 := make([]string, 3)
	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"
	fmt.Println(s2)

	// 超过容量
	s2 = append(s2, "d")
	s2 = append(s2, "e")
	s2 = append(s2, "f")
	fmt.Println(s2)

	// len 返回切片的长度
	s3 := make([]string, len(s2))

	// 支持复制函数
	copy(s3, s2)
	fmt.Println(s3)

	// 支持子切片 [i:j]  0 <= i <= j < cap(), 注意左闭右开
	// i 不写默认0 j 不写默认是len
	fmt.Println(s3[1:2])

	// 切片字面量
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// 多维 nxm
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		row := make([]int, 3)
		twoD[i] = row
	}
	fmt.Println(twoD)
}
