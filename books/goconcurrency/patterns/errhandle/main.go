package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// wrap error
	type Result struct {
		Error error
		Resp  *http.Response
	}
	client := http.Client{
		Timeout: time.Second * 3,
	}
	checkStatus := func(done <-chan struct{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				response, err := client.Get(url)
				result := Result{Error: err, Resp: response}
				select {
				case <-done:
					return
				// error return wrapped in result
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan struct{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://www.google.com", "https://www.bad.com"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Resp: %v\n", result.Resp)
	}
}
