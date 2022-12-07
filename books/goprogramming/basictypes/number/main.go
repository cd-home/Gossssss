package main

import "fmt"

func main() {
	/*
		整数
			int8 int16 int32 int64
			uint8 uint16 uint32 uint64
			int uint 自适应对应机器CPU架构位数

			无符号数往往只有在位运算或其它特殊的运算场景才会使用;
			大多数情况, 采用有符号;

			unicode 字符：rune = uint32
			byte = uint8

			uintptr 容纳指针大小
		浮点数
			float32 float64

		复数
			complex64 complex128

		算术运算
			+ - * / 用于 整数、浮点数、复数； / 表示整除
			% 只能用于整数
		比较运算符号(基础类型都是可比较的)
			> >= < <= !=
		位运算 (只能用于整型)
			&  AND
			|  OR
			^  XOR : 同为1，取0
			&^ AND NOT：第二个操作数位置为1, 对应第一个操作数置0
			>> n 右移				 x / 2^n
			<< n 左移 n 必须是无符号数  x * 2^n
		() 括号可以提高优先级
	*/
	fmt.Printf("%08b\n", 5)
	fmt.Printf("%08b\n", 3)
	fmt.Printf("%08b\n", 5&^3)

	// 进制
	fmt.Printf("%08b\n", 10) // 2
	fmt.Printf("%#o\n", 10)  // 8
	fmt.Printf("%#x\n", 10)  // 16
}
