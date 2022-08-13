package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var closed = false
var data = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int)  {
	for {
		if closed {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		if _, ok := data[x]; ok {
			closed = true
		} else {
			data[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Println("sum:", sum)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	n1, _ := strconv.Atoi(os.Args[0])
	n2, _ := strconv.Atoi(os.Args[1])

	if n1 > n2 {
		return
	}
	rand.Seed(time.Now().UnixNano())

	out := make(chan int)
	in := make(chan int)
	go first(n1, n2, out)
	go second(in, out)
	third(in)
}
