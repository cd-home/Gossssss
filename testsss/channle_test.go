package testsss

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// ch must be have a "write" goroutine meanwhile a "read" goroutine
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		ch <- true
	}()
	t.Log(<-ch) 
}