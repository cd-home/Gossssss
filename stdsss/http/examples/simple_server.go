package http

import (
	"fmt"
	"net/http"
)

// HelloWorld Handler
func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World, %s", r.URL.String())
}

func SimpleServer() {
	http.HandleFunc("/", logging(HelloWorld))
	http.ListenAndServe(":8080", nil)
}