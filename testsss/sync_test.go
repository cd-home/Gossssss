package testsss

import (
	"testing"
	"sync"
	"time"
)

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	testMap := make(map[int]int)
	lock := &sync.Mutex{}

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(j int) {
			defer wg.Done()
			lock.Lock()
			testMap[j] = j
			time.Sleep(time.Second * time.Duration(j))
			t.Log(testMap)
			lock.Unlock()
		}(i)
	}
	wg.Wait()
}