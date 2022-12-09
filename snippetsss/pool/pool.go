package pool

import (
	"context"
	"errors"
	"log"
	"sync"
)

type gPool[task Task] struct {
	mu       sync.Mutex
	task     []chan task
	number   int
	size     int
	cut      int
	state    []bool
	listener chan struct{}
	clear    context.CancelFunc
	rootCtx  context.Context
	close    bool
	catch    RuntimeErr
}

func New[task Task](ctx context.Context, n, s int) *gPool[task] {
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	pool := &gPool[task]{
		number:   n,
		size:     s,
		task:     make([]chan task, n),
		state:    make([]bool, n),
		listener: make(chan struct{}),
		rootCtx:  cancelCtx,
		clear:    cancelFunc,
	}
	for i := 0; i < n; i++ {
		pool.task[i] = make(chan task, s)
	}
	return pool
}

func (p *gPool[T]) Run() {
	p.start()
	go func() {
		for {
			select {
			case <-p.listener:
				go p.start()
			case <-p.rootCtx.Done():
				for i := 0; i < len(p.task); i++ {
					close(p.task[i])
				}
				p.close = true
				return
			}
		}
	}()
}

func (p *gPool[T]) start() {
	for i := 0; i < p.number; i++ {
		p.mu.Lock()
		if !p.state[i] && !p.close {
			p.state[i] = true
			ctx, cancelFunc := context.WithCancel(p.rootCtx)
			go func(ctx context.Context, tasks chan T, id int) {
				defer func(id int) {
					cancelFunc()
					p.state[id] = false
					p.listener <- struct{}{}
				}(id)
				for {
					select {
					case task := <-tasks:
						task.Run(ctx)
					case <-ctx.Done():
						p.close = true
						return
					}
				}
			}(ctx, p.task[i], i)
		}
		p.mu.Unlock()
	}
}

func (p *gPool[T]) Put(task T) error {
	for {
		p.next()
		if p.close {
			return errors.New("gPool Closed")
		}
		if p.state[p.cut] {
			p.task[p.cut] <- task
			return nil
		}
	}
}

func (p *gPool[T]) next() {
	if p.cut == p.number-1 {
		p.cut = 0
	} else {
		p.cut++
		log.Printf("Find Next G [%d]", p.cut)
	}
}

func (p *gPool[T]) check() bool {
	for i := 0; i < p.number; i++ {
		if p.state[i] {
			return true
		}
	}
	return false
}

func (p *gPool[T]) Close() {
	p.clear()
}
