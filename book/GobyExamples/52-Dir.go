package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CheckErr(e error)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic(e)
}

func main() {
	err := os.Mkdir("p", 0755)
	CheckErr(err)
	defer os.RemoveAll("p")

	err = os.MkdirAll("p/c/file", 0755)
	CheckErr(err)

	c, err := ioutil.ReadDir("p")
	CheckErr(err)
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Mode())
	}

	os.Chdir("./Go Code Style")
	c, err = ioutil.ReadDir(".")
	CheckErr(err)
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Mode())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(info.Name())
		return nil
	})

	// 回到原来的位置
	os.Chdir("../")
}
