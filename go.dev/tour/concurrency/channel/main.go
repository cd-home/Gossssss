package main

import (
	"fmt"
	"time"
)

// SendAndReceiveChannel
// Channels are a typed conduit[带类型的管道] through
// which you can send and receive[发送与接收值] values with the channel operator, <-.
// By default[无缓冲通道], sends and receives block[阻塞] until the other side is ready.
// This allows goroutines to synchronize[同步] without explicit locks[互斥锁] or
// condition variables[条件变量].
func SendAndReceiveChannel() {
	sum := func(slices []int, res chan int) {
		sums := 0
		for _, v := range slices {
			sums += v
		}
		// mock io
		time.Sleep(time.Second)
		// send
		res <- sums
	}

	s := []int{7, 2, 8, -9, 4, 0}

	// channels must be created before use
	// 使用之前必须初始化
	ch := make(chan int)

	// distributing the work between two goroutines
	go sum(s[:3], ch)
	go sum(s[3:], ch)

	// blocked receive
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}

// BufferedChannel
// Sends to a buffered channel block only when the buffer is full. 发送阻塞当通道满
// Receives block when the buffer is empty. 接收者阻塞当通道空
func BufferedChannel() {
	// Buffer Channel
	ch := make(chan int, 10)

	// Not be blocked if channel buffer not full
	ch <- 1
	ch <- 2
	// if sender do not send data, should close channel, or deadlock
	close(ch)

	// Read Until Close
	for v := range ch {
		fmt.Println(v)
	}
}

func ReadChannelByForLoop() {
	// read channel
	cds := make(chan string, 3)
	cds <- "Hello World!"
	cds <- "Go"
	cds <- "Python"
	// sender need close if no values to send
	close(cds)
	for {
		// ok is false if there are no more values to receive and the channel is closed.
		// PS: ok is false 表示, 通道关闭, 无数据可消费
		v, ok := <-cds
		if !ok {
			break
		}
		fmt.Println(ok, v)
	}
}

func RangeAndCloseChannel() {
	// close channel and for-range
	Fibonacci := func(n int) <-chan int {
		x, y := 0, 1
		// closure: do not expose ch
		// only expose read channel
		ch := make(chan int)
		go func() {
			// sender close ch
			// when is necessary to close a channel
			defer close(ch)
			for i := 0; i < n; i++ {
				ch <- x
				x, y = y, x+y
			}
		}()
		return ch
	}
	// for-range until close channel
	for v := range Fibonacci(10) {
		fmt.Println(v)
	}
}

func main() {
	// Default NoBufferedChannel
	SendAndReceiveChannel()

	// buffer channel
	BufferedChannel()

	// Read Channel
	ReadChannelByForLoop()
}
