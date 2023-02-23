package main

import (
	"fmt"
	"log"
	"unsafe"
)

func MovePointer() {
	arr := [...]int{1, 2, 3, 4, 5}
	pointer := &arr[0]
	log.Println(*pointer)
	addr := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	for i := 0; i < len(arr)-1; i++ {
		pointer := (*int)(unsafe.Pointer(addr))
		log.Println(*pointer)
		addr = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	}
}

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

	MovePointer()
}
