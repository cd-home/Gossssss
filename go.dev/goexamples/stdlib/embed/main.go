package main

import (
	"embed"
	"fmt"
)

//go:embed test.txt
var fileString string

//go:embed tests/test.txt
//go:embed tests/*.txt
var folder embed.FS

func main() {
	fmt.Println(fileString)

	file, err := folder.ReadFile("tests/test.txt")
	if err != nil {
		return
	}
	fmt.Println(string(file))

	// 下面的文件
	dir, err := folder.ReadDir("tests")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, d := range dir {
		fmt.Println(d.Name())
	}
}
