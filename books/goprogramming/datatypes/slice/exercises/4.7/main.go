package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(Reverse2([]byte(s)))
}

func Reverse(bs []byte) string {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func Reverse2(bs []byte) string {
	l := len(bs)
	for i := 0; i < l/2; i++ {
		bs[i], bs[l-i-1] = bs[l-i-1], bs[i]
	}
	return string(bs)
}
