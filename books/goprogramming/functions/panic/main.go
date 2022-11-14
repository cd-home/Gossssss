package main

import "log"

func main() {
	// 一般而言，当panic异常发生时，程序会中断运行，
	// 并立即执行在该goroutine中被延迟的函数（defer机制）
	defer func() {
		log.Println("123")
	}()
	panic("1")
	panic("2")
	panic("3")

	// panic 一般用于严重的错误;
	// 对于大部分错误，我们应该使用Go提供的错误机制，而不是panic，尽量避免程序的崩溃。
	// 在健壮的程序中，任何可以预料到的错误，如不正确的输入、错误的配置或是失败的I/O操作都应该被优雅的处理
	// 最好的处理方式，就是使用Go的错误机制
}
