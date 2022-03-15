package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		 fmt.Fprint(rw, "Hello World!")
	})
	_ = http.ListenAndServe(":8999", nil)
	
}
