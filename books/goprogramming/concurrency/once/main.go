package main

import (
	"fmt"
	"sync"
)

var loadOnce sync.Once

func main() {
	LoadConfig()
	LoadConfig()
}

func LoadConfig() {
	// 只执行一次
	loadOnce.Do(func() {
		fmt.Println("load config")
	})
}
