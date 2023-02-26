package main

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pwd, _ := os.Getwd()
	path := pwd + "/go.dev/goexamples/stdlib/os/file/main.go"
	// 全部读取
	data, err := os.ReadFile(path)
	Check(err)
	fmt.Println(string(data))

	fmt.Println("==========================================")

	// 文件具柄 只读打开
	f, err := os.Open(path)
	Check(err)
	buf := make([]byte, 100)
	n, err := f.Read(buf)
	Check(err)
	fmt.Println(string(buf[0:n]))

	// 文件具柄  更多的文件操作、读、写、信息、权限 4 2 1
	//os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
}
