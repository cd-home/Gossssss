package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// md5 hash 函数, 对相同的输入返回一致结果
	// 常用与验证完整性, 例如文件完整性
	hash := md5.New()
	hash.Write([]byte("Hello Md5"))
	s := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(s)
	fmt.Println(len(s))
}
