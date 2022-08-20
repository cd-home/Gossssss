package main

import (
	"log"
	"unsafe"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	pointer := &arr[0]
	log.Println(*pointer)
	addr := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	for i := 0; i < len(arr) - 1; i++ {
		pointer := (*int)(unsafe.Pointer(addr))
		log.Println(*pointer)
		addr = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	}
}
