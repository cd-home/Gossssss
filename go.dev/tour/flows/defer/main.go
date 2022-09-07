package main

import (
	"fmt"
)

// DeferFlow
// A defer statement defers the execution of a function
// until the surrounding function returns.
// 在函数返回之前执行
func DeferFlow() {
	fmt.Println("Testing deferred")
	// reverse order output
	defer fmt.Println("Testing deferred")
}

// DeferStack
// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.
// 当 defer 调用 的时候, 会 push 进入栈
// 函数返回之前 依次出栈执行 [先进后出]
func DeferStack() {
	defer fmt.Println("Testing deferred 1")
	defer fmt.Println("Testing deferred 2")
	defer fmt.Println("Testing deferred 3")
}

func DeferredPanic() {
	// catch panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("TestDeferredPanic")
}

func main() {
	DeferFlow()
	DeferStack()
	DeferredPanic()
}
