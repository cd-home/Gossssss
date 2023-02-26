package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CheckErr(e error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic(e)
}

// directory

func main() {
	err := os.Mkdir("p", 0755)
	CheckErr(err)
	defer os.RemoveAll("p")

	err = os.MkdirAll("p/c/file", 0755)
	CheckErr(err)

	c, err := os.ReadDir("p")
	CheckErr(err)
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	// 改变当前的工作目录
	os.Chdir("./books")
	c, err = os.ReadDir(".")
	CheckErr(err)
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(d.Name())
		return nil
	})

	// 回到原来的位置
	os.Chdir("../")

	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "sampleDir")
	if err != nil {
		return
	}
	defer os.RemoveAll(tempDir)
}
