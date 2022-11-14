package main

import "fmt"

func main() {
	defer func() {
		// 捕获panic的信息, 从panic中恢复程序
		// 只恢复应该被恢复的panic异常，某些致命异常应该直接的退出
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("1")
}
