package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// MD5、SHA1 SHA256 (签名算法) 主要用于验证，防止信息被修改
	// 文件校验、数字签名、鉴权协议
	hash := sha256.New()
	hash.Write([]byte("Hello SHA256"))
	s := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(s)
	fmt.Println(len(s))
}
