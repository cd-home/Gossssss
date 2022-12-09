package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	// hash 函数 (摘要、签名算法)
	hash := sha1.New()
	hash.Write([]byte("Hello SHA1"))
	s := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(s)
	fmt.Println(len(s))

	// file md5
	hash = sha1.New()
	path, _ := os.Getwd()
	f, _ := os.Open(path + "/go.dev/stdlib/crypto/sha1/main.go")
	if _, err := io.Copy(hash, f); err != nil {
		fmt.Println(err)
		return
	}
	s = hex.EncodeToString(hash.Sum(nil))
	fmt.Println(s)
}
