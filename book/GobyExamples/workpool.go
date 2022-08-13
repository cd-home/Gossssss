package main

import (
	"sync"
	"time"
)

func main() {

}

var (
	s = 10
	clients = make(chan Client, s)
	resp = make(chan Resp, s)
)

type Client struct {
	id int
	integer int
}

type Resp struct {
	job Client
	square int
}

func worker(w *sync.WaitGroup) {
	for c := range clients {
		squre := c.integer * c.integer
		out := Resp{
			job:    c,
			square: squre,
		}
		resp <- out
		time.Sleep(time.Second)
	}
	wg.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&w)
	}
	wg.Wait()
	close(resp)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{
			id:      i,
			integer: i,
		}
		clients <- c
	}
	close(clients)
}