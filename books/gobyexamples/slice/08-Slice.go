package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)

	// 会初始化
	s1 := make([]int, 3, 3)
	fmt.Println(s1)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	fmt.Println(s1)

	s2 := make([]string, 3)
	s2[0]= "a"
	s2[1]= "b"
	s2[2]= "c"
	fmt.Println(s2)

	// 超过容量
	s2 = append(s2, "d")
	s2 = append(s2, "e")
	s2 = append(s2, "f")
	fmt.Println(s2)

	s3 := make([]string, len(s2))
	copy(s3, s2)
	fmt.Println(s3)

	fmt.Println(s3[1:2])
}
