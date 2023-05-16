package main

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func main() {
	Confinement()
	ConfinementGood()
}

func Confinement() {
	data := make([]int, 10)
	mock := func(data []int) {
		for i := 0; i < len(data); i++ {
			data[i] = i
		}
	}
	mock(data)
	// 在 loop 中 我只需要 写
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	// 但是这里暴露出来了
	handleData := make(chan int)
	go loopData(handleData)
	// 读
	for num := range handleData {
		fmt.Println(num)
	}
}

func ConfinementGood() {
	chanOwner := func() <-chan int {
		// 写操作约束在下面的闭包中, 防止其他g 写入
		results := make(chan int, 10)
		go func() {
			// 发送完成关闭
			defer close(results)
			for i := 0; i < 10; i++ {
				results <- i
			}
		}()
		// 返回一个只读的chan
		return results
	}

	// 参数约束为一个只读的
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Println(result)
		}
		fmt.Println("Done receiving!")
	}
	results := chanOwner()
	consumer(results)
}

func TestNotSafe(t *testing.T) {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buf bytes.Buffer
		buf.Write(data)
		fmt.Println(buf.String())
	}
	var wg sync.WaitGroup

	wg.Add(2)

	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
