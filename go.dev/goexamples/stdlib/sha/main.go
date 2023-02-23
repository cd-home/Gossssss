package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	data := "Hello Golang"
	h := sha1.New()
	h.Write([]byte(data))

	sh := h.Sum(nil)
	fmt.Println(sh)
	//%x 提供了十六进制编码
	fmt.Printf("%x\n", sh)

	md := md5.New()
	md.Write([]byte(data))

	mdd := md.Sum(nil)

	fmt.Printf("%x\n", mdd)
}
