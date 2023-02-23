package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// md5 hash 函数, 对相同的输入返回一致结果, 任意长度的字节数组转换成一个定长的整数，并且这种转换是不可逆的
	// 常用与验证完整性, 例如文件完整性、密码存储、或者做数据的摘要
	hash := md5.New()
	hash.Write([]byte("Hello Md5"))
	// Sum 只是追加额外的, 一般情况下不需要
	s := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(s)
	fmt.Println(len(s))
	// 一般输出 16进制字符串
	fmt.Printf("%x", hash.Sum(nil))
}
