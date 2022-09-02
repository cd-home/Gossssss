package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Job)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Job(w http.ResponseWriter, r *http.Request) {

	ct, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)

	defer cancelFunc()

	success := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(4 * time.Second)
		fmt.Fprintf(w, "Hello\n")
		success <- struct{}{}
	}()

	select {
	case <-success:
		return
	case <-ct.Done():
		err := ct.Err()
		fmt.Println("server timeout:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
