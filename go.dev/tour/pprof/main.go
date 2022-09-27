package main

import (
	"bytes"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		buffer := bytes.NewBufferString("Hello World")
		buffer.WriteTo(rw)
	})
	var m sync.Mutex
	var data = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			data[i] = struct{}{}
			m.Unlock()
		}(i)
	}
	log.Println("Listening :8080")
	http.ListenAndServe(":8080", nil)
}
