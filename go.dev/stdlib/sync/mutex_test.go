package sync_test

import (
	"fmt"
	"sync"
	"testing"
)

// go test -race -v mutex_test.go

type SafeMap struct {
	sync.Mutex
	m map[string]string
}

func (m *SafeMap) Set(key string, value string) {
	defer m.Unlock()
	m.Lock()
	m.m[key] = value
}

func TestSafeMap(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)

	sm := &SafeMap{m: make(map[string]string)}
	for i := 0; i < 10; i++ {
		go func(i int) {
			sm.Set(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
			wg.Done()
		}(i)
	}
	
	wg.Wait()
	t.Log(sm)
}
