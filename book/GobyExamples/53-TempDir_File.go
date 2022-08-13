package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := ioutil.TempFile("", "tmp.go")
	fmt.Println(f.Name())

	stat, _ := os.Stat(f.Name())
	fmt.Println(stat.ModTime())

	defer os.RemoveAll(f.Name())

}
