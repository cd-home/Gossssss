package main

import (
	"fmt"
	"log"
	"unsafe"
)

func main() {

	var value int64 = 10
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	*p2 = 114523

	log.Println(value)

	var test int
	test = 256
	var pp = (*uint8)(unsafe.Pointer(&test))
	*pp = 11
	fmt.Println(test)
}
