package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encodeString := base64.StdEncoding.EncodeToString([]byte("God Yao"))
	fmt.Println(encodeString)

	decodeString, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Println(string(decodeString))
}
