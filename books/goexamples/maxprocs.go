package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(100)
	fmt.Println(runtime.GOMAXPROCS(8))

	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	n := 10

	fmt.Printf("%d", n)
	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", <-createNumber)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Ex...")
	end <- true
}

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()")
		}
	}

}