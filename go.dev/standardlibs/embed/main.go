package main

import (
	"embed"
	_ "embed" // 读取单文件需要
	"fmt"
)

// 打包静态资源， 注意静态资源一旦被打包，那么只可读
// //go:embed static  目录匹配
// //go:embed static/* 模糊匹配
// //go:embed index.html home.html 多文件
//go:embed static/*
var f embed.FS

func main() {
	// 读文件
	data, _ := f.ReadFile("static/css/index.css")
	fmt.Println(string(data))

	// 读取目录
	dirEntries, _ := f.ReadDir("static")
	fmt.Println(dirEntries[0].IsDir())
}
