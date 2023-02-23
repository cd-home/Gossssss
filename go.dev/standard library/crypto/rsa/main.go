package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	// openssl genrsa -out rsa_private_key.pem 1024
	// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
	origin := "I need rsa encrypt."
	bytes, _ := RsaEncrypt(origin)

	// 解密
	ori, _ := RsaDecrypt(string(bytes))
	fmt.Println(string(ori))
}

func RsaEncrypt(origin string) ([]byte, error) {
	block, _ := pem.Decode(public)
	x, _ := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := x.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(origin))
}

func RsaDecrypt(cipher string) ([]byte, error) {
	block, _ := pem.Decode(private)
	private, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return rsa.DecryptPKCS1v15(rand.Reader, private, []byte(cipher))
}

var private = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDFY2SaK2P+BqE0DJ/kOqOqwwdV3Q2oxKyueyn8TMeH16s0Gwui
/JVv21SPEkfu7VndtoB6c6OZSQelXyLm96vocwwmMGN6v47ZcIntvvVzlJ3I/FIx
rQ7lUHP7JmJ8ky6XkjI/o9N7k1ieL0ZLDp2ZIAP5/SDWhEILXdZGvMkM1QIDAQAB
AoGAFyMZbcDcTbB1QOah71y4RpHp0DTDHx8+po0lVL+jO9cZ6FerO4Wj9qKc/NBF
wknsGMMFuFTJvnQDA21tZ+lsKvRJ/bjEBG67wka/rNh0BS7bx1gsTk0jXn8mqIhs
l1SWIvI5xt3jJEyWCSwnZQa6hmxp2k5SqZbOumLHLs/ZIpkCQQDx6R5u4rCvihDm
Fj2J1VQ6S5H0xW4XjRUb81bO5biy+DsvnVBjGGQ1f1nVBcNoRW5cKx/g2NhESfox
CtUjbEvTAkEA0OJyXpqtC4v7LWk85mJALoAhK41ZGNE0baYtEe9G3n9KccD0EBOY
aKJ9Q8W491aYgVJld+RJ6h03tvkcuFsjtwJABVKdys7OGG6vRIDExd2dxtKW2Y3m
iDogdsb55/+B+t4fX0LU2/lTayHsNhW0YPz2Gq0QPIBI8ee2y5Fzcx4U5QJBAIDn
/iSb18aUcEGp/EdAwtLvh4jVBfLmyDASqSW0QEv2yf5BUVzKeve3rw4v6uHYsuy7
6JqAa4zq0wM74Jl0xSkCQB8/GRJyrQGI5FRyxBbAob1EDKbAoZoZ+zvAB1GODehq
TbM8bRAESr7SezK5QMFpkE5cdLeWeav+lanD5lDvfXQ=
-----END RSA PRIVATE KEY-----`)

var public = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDFY2SaK2P+BqE0DJ/kOqOqwwdV
3Q2oxKyueyn8TMeH16s0Gwui/JVv21SPEkfu7VndtoB6c6OZSQelXyLm96vocwwm
MGN6v47ZcIntvvVzlJ3I/FIxrQ7lUHP7JmJ8ky6XkjI/o9N7k1ieL0ZLDp2ZIAP5
/SDWhEILXdZGvMkM1QIDAQAB
-----END PUBLIC KEY-----`)
