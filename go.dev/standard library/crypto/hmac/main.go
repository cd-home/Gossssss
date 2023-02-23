package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Package hmac implements the Keyed-Hash Message Authentication Code
	// 加盐的hash算法, hash算法, 亦可称为散列函数
	// 不能泄漏密钥
	secret := []byte("I am secret")
	hash := hmac.New(sha256.New, secret)

	// 传输数据
	data := []byte("I am data")

	hash.Write(data)
	// 加密数据
	s := hash.Sum(nil)
	fmt.Printf("%x", s)

	// 验证数据是否被篡改
	fmt.Println(ValidMAC(data, s, secret))
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
