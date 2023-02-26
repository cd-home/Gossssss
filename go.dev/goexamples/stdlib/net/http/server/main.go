package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello World!")
}
