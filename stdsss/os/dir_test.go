package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	// 创建
	os.Mkdir("dir", 0777)

	// 递归
	os.MkdirAll("dir/sub", 0777)
	os.MkdirAll("dir/sub2", 0777)

	// 要求目录必须是空的
	// os.Remove("dir")

	// 修改名称
	os.Rename("dir/sub", "dir/sub1")

	file, _ := os.Open("dir")

	// 参数一般设置大一点1024，这样可以获取到所有的，可以参考RemoveAll 实现
	fmt.Println(file.Readdirnames(3))

	// 	递归删除
	os.RemoveAll("dir")
}
