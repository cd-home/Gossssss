package sync_test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

type Test struct {
	Name string `json:"name"`
}

func WaitForAllGroutines(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}(i)
	}
	wg.Wait()
	tt := Test{Name: "Mike"}
	json.Marshal(tt)
}
