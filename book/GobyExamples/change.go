package main

import (
	"fmt"
	"time"
)

func main() {
	plus := make(chan int)
	reduce := make(chan int)
	go change(plus, reduce)

	for i := 1; i <= 101; i++ {
		go func(i int) {
			plus <- i
		}(i)
	}

	for i := 1; i <= 101; i++ {
		go func(i int) {
			reduce <- i
		}(i)
	}

	time.Sleep(time.Second * 5)
}

func change(plus chan int, reduce chan int) {
	num := 0
	for {
		select {
		case d1 := <-plus:
			if d1 == 101 {
				fmt.Println(num)
			} else {
				num += d1
			}
		case d2 := <-reduce:
			if d2 == 101 {
				fmt.Println(num)
			} else {
				num -= d2
			}
		}
	}
}
