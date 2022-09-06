package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGoProcessNums(t *testing.T) {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for char := 0; char < 26; char++ {
			fmt.Printf("%c", 'A'+char)
			// Gosched yields the processor, allowing other goroutines to run
			runtime.Gosched()
		}
	}()

	go func() {
		defer wg.Done()
		for char := 0; char < 26; char++ {
			fmt.Printf("%c", 'a'+char)
			runtime.Gosched()
		}
	}()

	wg.Wait()
	t.Log("done")
}
