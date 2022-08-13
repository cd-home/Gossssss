package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvironment(t *testing.T) {
	// 设置环境变量
	err := os.Setenv("Key", "Value")
	if err != nil {
		return
	}
	// 获取环境变量
	fmt.Println(os.Environ())
	os.Getenv("")
}
