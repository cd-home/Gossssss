package main

import (
	"fmt"
	"runtime"
	"time"
)

// SwitchCase
// A switch statement is a shorter way to write a sequence of if - else statements
// It runs the first case whose value is equal to the condition expression.
// 执行第一个符合条件的 case (从上到下), 如果都没有并且有default的情况就会执行default
// 注意, 并不需要显式的break, Go 自动提供了
// 如果有多个if-else语句, 可采用 switch-case
func SwitchCase() {
	fmt.Print("Go runs on ")
	// 支持变量声明
	switch os := runtime.GOOS; os {
	// case 是 可比较的即可
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough, 执行下一个case
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// SwitchStyle
// switch 可没有条件, 即是 case 可以是表达式, 求的结果是布尔值即可
func SwitchStyle() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	SwitchCase()
	SwitchStyle()
}
