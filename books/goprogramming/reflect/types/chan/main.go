package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int, 10)

	rv := reflect.ValueOf(&ch)
	elem := rv.Elem()
	// 获取容量长度
	fmt.Println(elem.Cap())
	fmt.Println(elem.Len())

	// 发送数据
	elem.Send(reflect.ValueOf(10))
	//elem.TrySend()
	// 接收数据
	fmt.Println(elem.Recv())
	fmt.Println(elem.TryRecv())
}
