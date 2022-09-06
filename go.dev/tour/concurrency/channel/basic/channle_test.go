package basic

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// channel must have a "write" goroutine meanwhile a "read" goroutine
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		ch <- true
	}()
	t.Log(<-ch)
}

func TestNoCacheChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		// write close
		close(ch)
	}()
	for v := range ch {
		t.Log(v)
	}
}
