package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Q struct {
	items []int
	size  int
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewQ(size int) *Q {
	q := &Q{
		items: make([]int, 0, size),
		size:  size,
	}
	q.cond = sync.NewCond(&q.lock)
	return q
}

func (q *Q) Send(i int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == q.size {
		q.cond.Wait()
	}
	q.items = append(q.items, i)
	log.Printf("send item: %d, len=%d\n", i, len(q.items))
	q.cond.Signal()
}

func (q *Q) Receive() {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		q.cond.Wait()
	}
	i := q.items[0]
	q.items = q.items[1:]
	log.Printf("receive item: %d, len=%d", i, len(q.items))
	q.cond.Signal()
}

func main() {
	q := NewQ(5)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++ {
		go func() {
			for {
				item := rand.Int()
				q.Send(item)
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}()
	}

	for i := 0; i < 2; i++ {
		go func() {
			for {
				q.Receive()
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}()
	}
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-sign
}
