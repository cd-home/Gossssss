package main

import "fmt"

func Ping(pings chan<- string, msg string) {
	pings <- msg
}

func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// 无方向通道可以转换为有方向的通道
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// 有方向的通道通常作为函数参数, 这样可以在函数内部做到读、写的安全性
	Ping(pings, "passed message")
	Pong(pings, pongs)

	fmt.Println(<-pongs)
}
