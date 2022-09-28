package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ch = make(chan int, 10)
	rvch := reflect.ValueOf(ch)
	fmt.Println(rvch.Len())
	fmt.Println(rvch.Cap())
	fmt.Println(rvch.Kind())
	// send
	rvch.Send(reflect.ValueOf(8))
	rvch.Send(reflect.ValueOf(9))
	fmt.Println(rvch.Len())
	fmt.Println(rvch.Recv())
	v, ok := rvch.TryRecv()
	if ok {
		fmt.Println(v)
	}
	rtch := reflect.TypeOf(ch)
	fmt.Println(rtch.Kind())

	fmt.Println(rtch.ChanDir())
}
