package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	// 对象池, 避免频繁的创建销毁对象(对象复用、减少内存分配)
	pool := &sync.Pool{New: func() any {
		log.Println("create new object")
		return "Hello"
	}}
	// 第一次没有的情况从New中创建
	obj := pool.Get().(string)
	fmt.Println(obj)

	// 使用完需要放回
	pool.Put(obj)

	obj2 := pool.Get().(string)
	fmt.Println(obj2)
}
