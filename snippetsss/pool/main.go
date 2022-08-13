package main

import (
	"context"
	"sync"
)

type TaskFace interface {
	Run(ctx context.Context)
}

type RuntimeErr interface {
	Catch()
}

type Pool[Task TaskFace] struct {
	mu       sync.Mutex
	task     []chan Task
	size     int
	number   int
	cut      int
	catch    RuntimeErr
	status   []bool
	listener chan struct{}
	clear    context.CancelFunc
	rootCtx  context.Context
	close    bool
}

func New[Task TaskFace](ctx context.Context, number int) *Pool[Task] {
	cancel, cancelFunc := context.WithCancel(ctx)
	return &Pool[Task]{
		number:   number,
		task:     make([]chan Task, number),
		status:   make([]bool, number),
		listener: make(chan struct{}),
		rootCtx:  cancel,
		clear:    cancelFunc,
	}
}

func (p *Pool[T]) Start() {
	p.start()
	go func() {
		for {
			select {
			case <-p.listener:
				go p.start()
			case <-p.rootCtx.Done():
				return
			}
		}
	}()
}

func (p *Pool[T]) start() {
	for i := 0; i < p.number; i++ {
		p.mu.Lock()
		if !p.status[i] && !p.close {
			p.status[i] = true
			ctx, _ := context.WithCancel(p.rootCtx)
			go func(ctx context.Context, tasks chan T, id int) {
				defer func(id int) {
					p.status[id] = false
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

func (p *Pool[T]) Add(task T) {
	p.next()
	for {
		if p.close {
			return
		}
		if !p.check() {

		}
		if p.status[p.cut] {
			p.task[p.cut] <- task
			return
		}
		p.next()
	}
}

func (p *Pool[T]) next() {
	if p.cut == p.number-1 {
		p.cut = 0
	} else {
		p.cut++
	}
}

func (p *Pool[T]) check() bool {
	for i := 0; i < p.number; i++ {
		if p.status[i] {
			return true
		}
	}
	return false
}

func (p *Pool[T]) Close() {
	p.clear()
}

func main() {

}
