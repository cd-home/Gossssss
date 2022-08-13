package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// 拼接
	path := filepath.Join("p", "c", "x_file.go")
	fmt.Println(path)

	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))

	fmt.Println(filepath.IsAbs(path))

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// 后缀
	fmt.Println(filepath.Ext("config.json"))
	fmt.Println(filepath.Ext("config.json")[1:])

	// 去除后缀
	fmt.Println(strings.TrimSuffix("config.json", ".json"))

	// 寻找相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rel)
}
