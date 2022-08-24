package main

import (
	"fmt"
	"time"
)

func Select() {
	Fibonacci := func(exit chan struct{}) <-chan int {
		ch := make(chan int)
		x, y := 0, 1
		go func() {
			defer close(ch)
			for {
				// case must be a channel
				// It chooses one at random if multiple are ready.
				select {
				case <-exit:
					fmt.Println("exit.")
					return
				case ch <- x:
					x, y = y, x+y
				}
			}
		}()
		return ch
	}

	exit := make(chan struct{})
	ch := Fibonacci(exit)
	for i := 0; i < 10; i++ {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		}
	}
	// receiver 10 times and close
	close(exit)
	time.Sleep(time.Second)
}

func BlockSelect() {
	// select blocked
	cjs := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 3)
		cjs <- 1
		close(cjs)
	}()
	select {
	case v := <-cjs:
		fmt.Println(v)
	}
}

func SelectDefault() {
	selectDefaultMode := func(done chan struct{}) {
		tick := time.Tick(time.Second * 2)
		after := time.After(time.Second * 4)
		go func() {
			for {
				select {
				case <-done:
					fmt.Println("SelectDefault exit.")
					return
				case t := <-tick:
					fmt.Println(t)
				case v := <-after:
					fmt.Println(v)
				default:
					fmt.Println("default....")
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
	}
	// select default
	kip := make(chan struct{})
	selectDefaultMode(kip)

	time.Sleep(time.Second * 8)
	close(kip)
}

func main() {
	Select()
	BlockSelect()
	SelectDefault()
	time.Sleep(2 * time.Second)
}
