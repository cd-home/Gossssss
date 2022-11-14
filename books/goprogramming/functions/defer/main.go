package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	// 延迟函数调用
	defer resp.Body.Close()

	f, err := os.Open("path")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// defer 对具名返回值有影响
	fmt.Println(double(2))
}

func double(x int) (res int) {
	// 被defer执行的匿名函数可以修改具名返回值
	defer func() {
		res++
	}()
	// defer语句中的函数会在return语句更新返回值变量后再执行
	// res = x * x
	// 然后执行defer, res++
	return x * x
}
