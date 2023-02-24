package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorld Handler
func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(rw, "Hello World, %s", r.URL.String())
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
