package main

import "fmt"

func main() {
	// 1000
	//fmt.Println(1000 << 1)

	fmt.Println(printNumbers(3))
}

func printNumbers(n int) []int {
	base := 10
	limit := 10
	for i := 1; i < n; i++ {
		limit *= base
	}
	fmt.Println(limit)
	var res []int
	for i := 1; i < limit; i++ {
		res = append(res, i)
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}
