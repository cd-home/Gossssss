package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "/Users/apple/Desktop/otherSpace/CodeExamples/Go by Example/46-URL.go"
	// 全部读取
	data, err := ioutil.ReadFile(path)
	Check(err)
	fmt.Println(string(data))

	fmt.Println("==========================================")

	// 文件具柄  更多的文件操作、读、写、信息、权限
	f, err := os.Open(path)
	Check(err)
	buf := make([]byte, 100)
	n, err := f.Read(buf)
	Check(err)
	fmt.Println(string(buf[0:n]))
}
