package main

import "fmt"

/*
1 操作数在左边，  右移就是除法  m >> n   m / 2**n  2的多少次方
2.              左移就是乘法  m << n   m * 2**n
*/
func main() {
	fmt.Println(1 << 10)
	fmt.Println(8 >> 1)
	fmt.Println(8 >> 2)
	fmt.Println(2 << 10) // 2 ** 2^10
	fmt.Println(67 << 1) // 67 * 2^1
}
