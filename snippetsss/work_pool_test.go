package snippetsss

import (
	"context"
	"errors"
	"log"
	"sync/atomic"
)

const Running = 1
const STOPPED = 0

// Task
type Task struct {
	Handler func(...interface{})
	Params  []interface{}
}

// Pool interface
type FacePool interface {
	Run()
	Put(*Task) error
	Close()
	TimeWaitDone(context.Context)
}

// goroutine Pool
type Pool struct {
	Cap          int64 // g的数量
	RunningWorks int64
	State        int
	TaskC        chan *Task // 任务队列
	CloseWorkers chan bool
	PanicFunc    func(interface{})
}

// NewPool
func NewPool(workers int64, tasks int) FacePool {
	return &Pool{
		Cap:          workers,
		State:        Running,
		TaskC:        make(chan *Task, tasks),
		CloseWorkers: make(chan bool),
		PanicFunc: func(i interface{}) {
			log.Printf("[Panic Func] %s", i.(error))
		},
	}
}

// Run
func (p *Pool) Run() {

	p.IncRunningWorks()

	go func() {

		defer func() {
			p.DecRunningWorks()
			if err := recover(); err != nil {
				p.PanicFunc(err)
			}
		}()

		for {
			select {
			case task, ok := <-p.TaskC:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			// Close Pool all goroutines
			case <-p.CloseWorkers:
				return
			}
		}
	}()
}

// Put
var ErrClosePool = errors.New("pool Stopped")

func (p *Pool) Put(task *Task) error {
	if p.State == STOPPED {
		return ErrClosePool
	}
	// 创建的g数量小于限制的g的数量
	if p.GetRunningWorks() < p.GetCap() {
		p.Run()
	}
	p.TaskC <- task
	return nil
}

// Close
func (p *Pool) Close() {
	p.State = STOPPED
	for len(p.TaskC) > 0 {
	}
	p.CloseWorkers <- true
	close(p.TaskC)
}

// WaitDone
func (p *Pool) TimeWaitDone(ctx context.Context) {
	// 超时
	for range ctx.Done() {
		p.State = STOPPED
		p.CloseWorkers <- true
		close(p.TaskC)
		return
	}

}

func (p *Pool) IncRunningWorks() {
	atomic.AddInt64(&p.RunningWorks, 1)
}

func (p *Pool) DecRunningWorks() {
	atomic.AddInt64(&p.RunningWorks, -1)
}

func (p *Pool) GetRunningWorks() int64 {
	return atomic.LoadInt64(&p.RunningWorks)
}

func (p *Pool) GetCap() int64 {
	return atomic.LoadInt64(&p.Cap)
}
